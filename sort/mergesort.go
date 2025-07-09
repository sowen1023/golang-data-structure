package sort

import "io.github.sowen.datastructure/util"

// MergeSort 自顶向下归并排序
func MergeSort(arr []int) {
	tmp := make([]int, len(arr), len(arr))
	util.RangeCopy(arr, 0, len(arr)-1, tmp)
	mergeSort(arr, 0, len(arr)-1, tmp)
}

func mergeSort(arr []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	m := (r-l)/2 + l

	mergeSort(arr, l, m, tmp)
	mergeSort(arr, m+1, r, tmp)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r, tmp)
	}
}

func MergeSort2(arr []int) {
	tmp := make([]int, len(arr), len(arr))
	util.RangeCopy(arr, 0, len(arr)-1, tmp)
	mergeSort2(arr, 0, len(arr)-1, tmp)
}

// 排序个数 16 以内使用插入排序以提高性能
func mergeSort2(arr []int, l, r int, tmp []int) {
	if l >= r {
		return
	}

	if r-l+1 <= 16 {
		for i := l; i <= r; i++ {
			t := arr[i]
			j := i
			for ; j-1 >= l && t < arr[j-1]; j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = t
		}
		return
	}

	m := (r-l)/2 + l

	mergeSort2(arr, l, m, tmp)
	mergeSort2(arr, m+1, r, tmp)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r, tmp)
	}
}

func merge(arr []int, l, m, r int, tmp []int) {
	util.RangeCopy(arr, l, r, tmp)

	i, j := l, m+1
	for k := l; i <= m || j <= r; k++ {
		if j > r || (i <= m && tmp[i] < tmp[j]) {
			arr[k] = tmp[i]
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
	//for k := l; i <= m || j <= r; k++ {
	//	if i > m {
	//		arr[k] = tmp[j]
	//		j++
	//	} else if j > r {
	//		arr[k] = tmp[i]
	//		i++
	//	} else if tmp[i] < tmp[j] {
	//		arr[k] = tmp[i]
	//		i++
	//	} else {
	//		arr[k] = tmp[j]
	//		j++
	//	}
	//}
}
