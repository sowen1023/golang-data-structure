package sort

import (
	"io.github.sowen.datastructure/util"
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	data := util.RandomIntSlices(10000, 100000, 5000000)
	copiedData := make([][]int, len(data), cap(data))
	for i, datum := range data {
		copiedItem := make([]int, len(datum), cap(datum))
		for j, v := range datum {
			copiedItem[j] = v
		}
		copiedData[i] = copiedItem
	}

	util.TestSort(t, ShellSort, data...)
	util.TestSort(t, QuickSort3Ways, copiedData...)
	if !reflect.DeepEqual(data, copiedData) {
		t.Fatal("invalid sort implementation")
	}
}
