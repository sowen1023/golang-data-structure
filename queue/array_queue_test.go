package queue_test

import (
	"fmt"
	"testing"

	"io.github.sowen.datastructure/queue"
)

func TestArrayQueue_Enqueue(t *testing.T) {
	queue := queue.NewArrayQueueWithCapacity[int](8)

	for i := 1; i <= 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	fmt.Printf("Head element = %v\n", queue.Head())
	fmt.Printf("Is empty = %v\n", queue.IsEmpty())

	for i := 1; i <= 10; i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

	fmt.Printf("Is empty = %v\n", queue.IsEmpty())
}
