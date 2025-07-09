package main

import (
	"fmt"

	"io.github.sowen.datastructure/list"
)

const N int = 10

func applyFn(fn func(int)) {
	for i := 0; i < N; i++ {
		fn(i)
	}
}

func main() {
	arr := list.NewWithCapacity[int](4)
	applyFn(func(i int) {
		arr.AddLast(i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.RemoveFirst()
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.AddFirst(i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.RemoveLast()
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.Add(arr.Size()/2, i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.Remove(arr.Size() / 2)
		fmt.Println(arr)
	})
}
