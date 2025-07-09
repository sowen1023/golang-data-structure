package sort

import (
	"io.github.sowen.datastructure/util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := [][]int{
		util.RandomIntSliceWithBound(10000, 10000),
		util.RandomIntSliceWithBound(100000, 100000),
		util.RandomIntSliceWithBound(1000000, 1000000),
		util.RandomIntSliceWithBound(2000000, 2000000),
		util.RandomIntSliceWithBound(5000000, 5000000),
	}

	copiedData := make([][]int, len(data), cap(data))
	for i, datum := range data {
		copiedItem := make([]int, len(datum), cap(datum))
		for j, v := range datum {
			copiedItem[j] = v
		}
		copiedData[i] = copiedItem
	}

	util.TestSort(t, MergeSort, data...)
	util.TestSort(t, MergeSort2, copiedData...)

	data = make([][]int, 1)
	data[0] = util.OrderedIntSlice(5000000, 5000000)
	util.TestSort(t, MergeSort2, data...)
}
