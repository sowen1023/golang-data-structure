package util

import (
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"io.github.sowen.datastructure/types"
)

func RandomIntSlice(length int) []int {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(math.MaxInt32)
	}
	return res
}

func RandomIntSliceWithBound(length, bound int) []int {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(bound)
	}
	return res
}

func ReversedIntSlice(length int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = length - i - 1
	}
	return res
}

func OrderedIntSlice(length, bound int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = i
	}
	return res
}

func RandomIntSlices(sizes ...int) [][]int {
	data := make([][]int, len(sizes))
	for i, size := range sizes {
		data[i] = RandomIntSliceWithBound(size, size)
	}
	return data
}

func VerifyIsSorted[T types.Comparable](arr []T) {
	var prev T
	for i, v := range arr {
		if i > 0 && v < prev {
			panic("Not a sorted slice")
		}
		prev = v
	}
}

func ResolveFuncName(fn interface{}) string {
	res := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	index := strings.LastIndex(res, ".") + 1
	res = string([]rune(res)[index:])
	return res
}

type Sorter[E types.Comparable] func([]E)

func (s Sorter[E]) Name() string {
	return ResolveFuncName(s)
}

func TestSort[E types.Comparable](t *testing.T, sorter Sorter[E], data ...[]E) {
	name := sorter.Name()
	for _, arr := range data {
		t.Logf("sort slice(length %9d) using %16v", len(arr), name)
		start := time.Now()
		sorter(arr)
		cost := time.Since(start)

		t.Logf("sort slice(length %9d) using %16v cost %v", len(arr), name, cost)
		VerifyIsSorted(arr)
	}
}

func TestSortWithName[E types.Comparable](t *testing.T, name string, sorter func([]E), data ...[]E) {
	for _, arr := range data {
		t.Logf("sort slice(length %9d) using %16v", len(arr), name)
		start := time.Now()
		sorter(arr)
		cost := time.Since(start)
		t.Logf("sort slice(length %9d) using %16v cost %v", len(arr), name, cost)
		VerifyIsSorted(arr)
	}
}

func Swap[E any](arr []E, i, j int) {
	// t := arr[i]
	// arr[i] = arr[j]
	// arr[j] = t
	arr[i], arr[j] = arr[j], arr[i]
}

func Copy[E any](src []E) []E {
	dst := make([]E, len(src))
	RangeCopy(src, 0, len(src)-1, dst)
	return dst
}

func RangeCopy[E any](src []E, l, r int, dest []E) {
	ArrayCopy(src, l, dest, l, r-l+1)
}

func ArrayCopy[E any](src []E, srcPos int, dest []E, destPos, length int) {
	for i, j := srcPos, 0; i < srcPos+length; i++ {
		dest[destPos+j] = src[i]
		j++
	}
}

func RandomStringSliceWithBound(length, bound int) []string {
	rand.Seed(time.Now().UnixNano())
	res := make([]string, length)
	for i := 0; i < length; i++ {
		sb := strings.Builder{}
		for j := 0; j < bound; j++ {
			// 33-126
			sb.WriteRune(rune(rand.Intn(94) + 33))
		}
		res[i] = sb.String()
	}
	return res
}
