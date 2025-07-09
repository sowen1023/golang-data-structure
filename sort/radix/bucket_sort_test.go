package radix

import (
	"io.github.sowen.datastructure/sort"
	"io.github.sowen.datastructure/util"
	"testing"
)

func TestBucketSort(t *testing.T) {
	// arr1 := util.RandomIntSliceWithBound(100, 1000)
	arr1 := util.RandomIntSlice(1_000_000)
	arr2 := util.Copy(arr1)
	arr3 := util.Copy(arr1)
	arr4 := util.Copy(arr1)

	util.TestSortWithName(t, "BucketSort", func(arr []int) {
		BucketSort(arr, 100)
	}, arr1)

	util.TestSortWithName(t, "BucketSort", func(arr []int) {
		BucketSort(arr, 11)
	}, arr2)

	util.TestSortWithName(t, "BucketSort2", func(arr []int) {
		BucketSort2(arr, 100)
	}, arr3)

	util.TestSort(t, sort.QuickSort2Ways, arr4)
}
