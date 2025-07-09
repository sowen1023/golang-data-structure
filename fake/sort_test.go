package fake

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"io.github.sowen.datastructure/util"
)

func QuickSort3Ways(arr []int) {
	quickSort3Ways(arr, 0, len(arr)-1)
	internalPrint("sort completed", arr, 0, len(arr)-1)
}

func quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	lt, gt := partition3Ways(arr, l, r)

	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}

// util invariant
// ① arr[l] 标定点元素
// ② arr[l+1, lt] < v && arr[lt+1, i-1] == v && [gt, r] > v
// ③ [i, gt-1] 待处理的区间
func partition3Ways(arr []int, l, r int) (int, int) {
	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	v := arr[l]

	// 初始变量阶段
	// ① arr[l+1, lt] 为 [l+1, l] 左闭 > 右闭的空区间，满足循环不变量
	// ② arr[gt,   r] 为 [r+1, r] 左闭 > 右闭的空区间，满足循环不变量
	lt, gt := l, r+1

	for i := lt + 1; i < gt; {
		if arr[i] < v {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > v {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}

	// 循环之后
	// ① [l+1, lt] < v
	// ② [lt + 1, gt - 1] == v
	// ③ [gt, r] > v

	arr[l], arr[lt] = arr[lt], arr[l]
	// ① arr[l, lt-1] < v
	// ② arr[lt, gt-1] == v
	// ③ [gt, r] > v
	return lt, gt
}

func QuickSort2Ways(arr []int) {
	quickSort2Ways(arr, 0, len(arr)-1)
	internalPrint("sort completed", arr, 0, len(arr)-1)
}

func quickSort2Ways(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition2Ways(arr, l, r)
	quickSort2Ways(arr, l, p-1)
	quickSort2Ways(arr, p+1, r)
}

// util invariant
// ① arr[l] 标定点元素
// ② arr[l+1, i-1] < v && [j+1, r] > v
// ③ i <= j
//
// 初始阶段
// i = l, j = r
// arr[l+1, l] 空区间满足 arr[l+1, i-1]
// arr[r+1, r]空区间满足 [j+1, r] > v
func partition2Ways(arr []int, l, r int) int {
	internalPrint("start partition", arr, l, r)

	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	internalPrint("after random", arr, l, r)

	v := arr[l]
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

		internalPrint("after swap", arr, l, r)
	}
	arr[l], arr[j] = arr[j], arr[l]

	internalPrint("after partition", arr, l, r)
	return j
}

func MergeSort(arr []int) {
	temp := make([]int, len(arr))
	util.RangeCopy(arr, 0, len(arr)-1, temp)
	mergeSort(arr, 0, len(arr)-1, temp)
}

func mergeSort(arr []int, l, r int, temp []int) {
	if l >= r {
		return
	}

	m := (r-l)/2 + l
	mergeSort(arr, l, m, temp)
	mergeSort(arr, m+1, r, temp)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r, temp)
	}
}

// 循环不变量
// ① l <= i <= m
// ② m < j <= r
// ③ i <= k <= r
// ④ arr[k - 1] <= arr[k]
func merge(arr []int, l, m, r int, temp []int) {
	util.RangeCopy(arr, l, r, temp)
	i, j, k := l, m+1, l
	for ; i <= m || j <= r; k++ {
		if j > r || (i <= m && temp[i] <= temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
	}
}

func BubbleSort(arr []int) {
	internalPrintAll("before sort", arr)

	sorted := false

	for i := 0; i < len(arr) && !sorted; i++ {
		sorted = true
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				sorted = false
			}
		}
		internalPrintAll("after util", arr)
	}
}

func SelectSort(arr []int) {
	internalPrintAll("before sort", arr)
	for i := 0; i < len(arr); i++ {
		p := i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[p] {
				p = j
			}
		}
		arr[i], arr[p] = arr[p], arr[i]
		internalPrintAll("after util", arr)
	}
}

// 循环不变量
// arr[0, i-1] 是有序的，将arr[i] 向 arr[0, i-1] 中插入，使 arr[0, i] 有序
func InsertSort(arr []int) {
	internalPrintAll("before sort", arr)

	for i := 1; i < len(arr); i++ {
		// 已排好序的
		internalPrint("sorted range", arr, 0, i-1)

		v, p := arr[i], i
		for ; p > 0 && v < arr[p-1]; p-- {
			arr[p] = arr[p-1]
		}
		// 找到 p 在排序后的数组中的位置
		arr[p] = v

		internalPrintAll("after util", arr)
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
	internalPrint("sort completed", arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// util invariant
// [l+1, p] < v && [p + 1, i-1] > v
// [i, r] 待处理的区间
func partition(arr []int, l, r int) int {
	internalPrint("start partition", arr, l, r)

	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	internalPrint("after random", arr, l, r)

	v, p := arr[l], l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			p++
			arr[p], arr[i] = arr[i], arr[p]
		}
	}
	arr[l], arr[p] = arr[p], arr[l]

	internalPrint("after partition", arr, l, r)
	return p
}

func internalPrintAll(msg string, arr []int) {
	internalPrint(msg, arr, 0, len(arr)-1)
}

func internalPrint(msg string, arr []int, l, r int) {
	temp := make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		if i < l || i > r {
			temp[i] = "❤️"
		} else {
			temp[i] = arr[i]
		}
	}
	fmt.Printf("%16v = %v\n", msg, temp)
}

func TestSorts(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	data := util.RandomIntSliceWithBound(20, 100)

	util.TestSort(t, BubbleSort, util.Copy(data))
	util.TestSort(t, SelectSort, util.Copy(data))
	util.TestSort(t, InsertSort, util.Copy(data))

	util.TestSort(t, MergeSort, util.Copy(data))

	util.TestSort(t, QuickSort, util.Copy(data))
	util.TestSort(t, QuickSort2Ways, util.Copy(data))
	util.TestSort(t, QuickSort3Ways, util.Copy(data))
}
