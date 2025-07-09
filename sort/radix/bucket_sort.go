package radix

import (
	"io.github.sowen.datastructure/sort"
	"math"
)

func BucketSort(arr []int, B int) {
	if B <= 1 {
		panic("B must > 1")
	}

	// 临时的额外空间
	temp := make([]int, len(arr))
	bucketSort(arr, 0, len(arr)-1, B, temp)
}

// bucketSort 对 arr[left, right] 区间进行桶排序
func bucketSort(arr []int, left int, right int, B int, temp []int) {
	// 递归到底的情况
	if left >= right {
		return
	}

	min, max := math.MaxInt32, math.MinInt32
	for i := left; i <= right; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 最大值与最小值相等表示所有元素全相等
	if max == min {
		return
	}

	// 每个桶里放多少个元素
	d := (max - min + 1) / B
	if (max-min+1)%B > 0 {
		d += 1
	}

	cnt := make([]int, B)
	index := make([]int, B+1)

	// O(n)
	// (arr[i]-min)/d 表示 arr[i]第几个桶
	for i := left; i <= right; i++ {
		cnt[(arr[i]-min)/d]++
	}

	// O(R)
	for i := 0; i < B; i++ {
		index[i+1] = index[i] + cnt[i]
	}

	// O(R)
	for i := left; i <= right; i++ {
		p := (arr[i] - min) / d
		temp[left+index[p]] = arr[i]
		index[p]++
	}

	// O(R)
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}

	// O(R)
	// 对第 0 个桶里的元素进行排序
	bucketSort(arr, left, left+index[0]-1, B, temp)

	// 从第二个桶开始，对后面所有的桶内元素进行排序
	for i := 0; i < B-1; i++ {
		bucketSort(arr, left+index[i], left+index[i+1]-1, B, temp)
	}
}

// BucketSort2 c 表示每个桶里多少个元素
func BucketSort2(arr []int, c int) {
	// 递归到底的情况
	if c <= 0 {
		panic("c must > 0")
	}

	min, max := math.MaxInt32, math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	// arr 中的数据范围
	rng := max - min + 1

	// 根据数据范围决定桶的个数
	B := rng / c
	if rng%c > 0 {
		B += 1
	}

	buckets := make([][]int, B)

	for i := 0; i < len(arr); i++ {
		p := (arr[i] - min) / c
		buckets[p] = append(buckets[p], arr[i])
	}

	for i := 0; i < B; i++ {
		sort.BubbleSort(buckets[i])
	}

	var index int
	for i := 0; i < B; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
}
