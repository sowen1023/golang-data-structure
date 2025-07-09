package sort

import (
	"io.github.sowen.datastructure/util"
	"math"
)

// MergeSortBU 自底向上归并排序
func MergeSortBU(arr []int) {
	tmp := make([]int, len(arr), len(arr))
	util.RangeCopy(arr, 0, len(arr)-1, tmp)

	n := len(arr)
	for sz := 1; sz < n; sz += sz {
		for i := 0; i+sz < n; i += sz * 2 {
			if arr[i+sz-1] > arr[i+sz] {
				mergeBU(arr, i, i+sz-1, int(math.Min(float64(i+sz+sz-1), float64(n-1))), tmp)
			}
		}
	}
}

func mergeBU(arr []int, l, m, r int, tmp []int) {
	util.RangeCopy(arr, l, r, tmp)

	i, j := l, m+1
	for k := l; i <= m || j <= r; k++ {
		if i > m {
			arr[k] = tmp[j]
			j++
		} else if j > r {
			arr[k] = tmp[i]
			i++
		} else if tmp[i] < tmp[j] {
			arr[k] = tmp[i]
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
}
