package sort

import (
	"testing"

	"io.github.sowen.datastructure/util"
)

func TestBubbleSort(t *testing.T) {
	util.TestSort(t, BubbleSort, util.RandomIntSlice(10000))
}
