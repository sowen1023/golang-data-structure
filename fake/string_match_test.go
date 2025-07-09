package fake

import (
	"fmt"
	"testing"

	"io.github.sowen.datastructure/util"
)

func KMP(s, t string) int {
	src, target := []rune(s), []rune(t)
	lps := GetLPS(string(target))

	var i, j int
	for i+len(target) <= len(src) {
		if src[i] == target[j] {
			i++
			j++
			if j == len(target) {
				return i - j
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}
	return -1
}

func GetLPS(src string) []int {
	runes := []rune(src)
	lps := make([]int, len(runes))

	for i := 1; i < len(runes); i++ {
		a := lps[i-1]
		for a > 0 && runes[i] != runes[a] {
			a = lps[a-1]
		}
		if runes[i] == runes[a] {
			lps[i] = a + 1
		}
	}

	return lps
}

func TestMatch(t *testing.T) {
	src := "owen欧文雪儿❤"
	targets := []string{
		"o", "w", "e", "n", "欧文雪", "Owen", "文", "❤",
	}
	fns := []func(string, string) int{KMP}

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
