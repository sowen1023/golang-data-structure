package sort

import (
	"io.github.sowen.datastructure/util"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	util.TestSort(t, InsertionSort, util.RandomIntSlices(10000, 100000, 200000)...)
}

func TestInsertionSort2(t *testing.T) {
	util.TestSort(t, InsertionSort2, util.RandomIntSlices(10000, 100000, 200000)...)
}
