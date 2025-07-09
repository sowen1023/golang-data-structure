package sort

import (
	"io.github.sowen.datastructure/util"
	"testing"
)

func selectSortData() [][]int {
	return [][]int{
		util.RandomIntSliceWithBound(100000, 100000),
		// RandomIntSliceWithBound(1000000, 1000000),
		// RandomIntSliceWithBound(2000000, 2000000),
		// RandomIntSliceWithBound(5000000, 5000000),
		// RandomIntSliceWithBound(10000000, 10000000),
		// RandomIntSlice(10000000),
		// OrderedIntSlice(10000000, 10000000),
		// RandomIntSliceWithBound(10000000, 1),
	}
}

func TestSelectSort(t *testing.T) {
	data := selectSortData()
	util.TestSort(t, SelectSort, data...)
}
