package match

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"io.github.sowen.datastructure/util"
)

//	func TestRandomFloat(t *testing.T) {
//		// [1.23 - 4.78)
//		for i := 0; i < 100; i++ {
//			fmt.Println(rand.Float64()*(4.78-1.23) + 1.23)
//		}
//	}

func TestMatch(t *testing.T) {
	src := "owen欧文雪儿❤"
	targets := []string{
		"o", "w", "e", "n", "欧文雪", "Owen",
	}
	fns := []func(string, string) int{BruteForce, RabinKarp, KMP, KMP2, KMPUtf8Bad, strings.Index}

	for _, target := range targets {
		for _, fn := range fns {
			index := fn(src, target)
			if index >= 0 {
				fmt.Printf("found [%v] at index [%v] use [%10v]\n", target, index, util.ResolveFuncName(fn))
			} else {
				fmt.Printf("can't find [%v]\n", target)
			}
		}
	}

	// 19 ??
	fmt.Printf("length of %v is %v\n", src, len(src))
}

func TestMatch2(t *testing.T) {
	var sb strings.Builder
	sb.WriteString("欧文owen欧文雪儿❤")
	sb.WriteString("[欧文欧文雪儿]")
	sb.WriteString("欧文雪儿❤雪儿欧文欧州雪儿❤")
	sb.WriteString("雪儿❤")

	src := sb.String()

	targets := []string{
		"欧文雪儿❤雪儿欧文欧州雪儿❤", "❤",
	}
	fns := []func(string, string) int{BruteForce, RabinKarp, KMP, KMP2, KMPUtf8Bad, strings.Index}

	for _, target := range targets {
		for _, fn := range fns {
			index := fn(src, target)
			if index >= 0 {
				fmt.Printf("found [%v] at index [%v] use [%10v]\n", target, index, util.ResolveFuncName(fn))
			} else {
				fmt.Printf("can't find [%v]\n", target)
			}
		}
	}
}

func testMatchPerformance(src, target *string, fn func(string, string) int) {
	start := time.Now()
	index := fn(*src, *target)
	fmt.Printf("found target string at index [%v] use [%10v] cost [%v]\n", index, util.ResolveFuncName(fn), time.Since(start))
}

func TestMatchPerformance1(t *testing.T) {
	src := util.GetPrideAndPrejudiceAsString()
	target := "china"
	testMatchPerformance(&src, &target, BruteForce)
	testMatchPerformance(&src, &target, RabinKarp)
	testMatchPerformance(&src, &target, KMP)
	testMatchPerformance(&src, &target, KMP2)

	// 处理中文返回的不是字符在字符串中的下标
	testMatchPerformance(&src, &target, KMPUtf8Bad)
	testMatchPerformance(&src, &target, strings.Index)
}

func TestMatchPerformance2(t *testing.T) {
	n, m := 1_000_000, 10_000
	var sBuilder strings.Builder
	for i := 0; i < n; i++ {
		sBuilder.WriteRune('a')
	}

	var tBuilder strings.Builder
	for i := 0; i < m-1; i++ {
		tBuilder.WriteRune('a')
	}
	tBuilder.WriteRune('b')

	src := sBuilder.String()
	target := tBuilder.String()

	testMatchPerformance(&src, &target, BruteForce)
	testMatchPerformance(&src, &target, RabinKarp)
	testMatchPerformance(&src, &target, KMP)
	testMatchPerformance(&src, &target, KMP2)

	testMatchPerformance(&src, &target, KMPUtf8Bad)
	testMatchPerformance(&src, &target, strings.Index)

}
