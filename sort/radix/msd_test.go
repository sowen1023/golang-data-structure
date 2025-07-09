package radix

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"

	"io.github.sowen.datastructure/util"
)

func TestMSDSort(t *testing.T) {
	arr := []string{"BCA", "CBAA", "AC", "BADFE", "ABC", "CBA"}
	MSDSort(arr)
	for _, s := range arr {
		fmt.Println(s)
	}
}

func DoSortTests(sort func([]string), arr []string) {
	sortFnName := runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name()

	log.Printf("sort slice(length %9d) using %16v\n", len(arr), sortFnName)
	start := time.Now()
	sort(arr)
	cost := time.Since(start)

	log.Printf("sort slice(length %9d) using %16v cost %v\n", len(arr), sortFnName, cost)
	util.VerifyIsSorted(arr)
}

func QuickSort2Ways(arr []string) {
	quickSort2ways(arr, 0, len(arr)-1)
}

func quickSort2ways(arr []string, l, r int) {
	if l >= r {
		return
	}

	p := partition2ways(arr, l, r)
	quickSort2ways(arr, l, p-1)
	quickSort2ways(arr, p+1, r)
}

func partition2ways(arr []string, l, r int) int {
	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	v := arr[l]
	// 循环不变量：arr[l + 1...i-1] < v && arr[j+1...r] > v
	i, j := l+1, r
	for {
		for i <= j && arr[i] < v {
			i++
		}
		for j >= i && arr[j] > v {
			j--
		}
		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]

	return j
}

func TestMSDSortPerformance(t *testing.T) {
	n, bound := 500000, 20

	arr := util.RandomStringSliceWithBound(n, bound)
	arr2 := make([]string, len(arr))

	copy(arr2, arr)

	DoSortTests(MSDSort, arr)
	DoSortTests(QuickSort2Ways, arr2)
}
