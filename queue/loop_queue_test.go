package queue_test

import (
	"fmt"
	"testing"

	"io.github.sowen.datastructure/queue"
)

func TestLoopQueue_Enqueue(t *testing.T) {
	q := queue.NewLoopQueueWithCapacity[int](8)

	n := 16
	for i := 0; i < n; i++ {
		q.Enqueue(i)
		fmt.Println(q)
	}
	fmt.Printf("Head = %v, empty = %v\n", q.Head(), q.IsEmpty())

	for i := 0; i < n; i++ {
		q.Dequeue()
		fmt.Println(q)
	}

	for i := 0; i < n; i++ {
		q.Enqueue(i)
		fmt.Println(q)
	}
	fmt.Printf("Head = %v, empty = %v\n", q.Head(), q.IsEmpty())

	for i := 0; i < n; i++ {
		q.Dequeue()
		fmt.Println(q)
	}
}
