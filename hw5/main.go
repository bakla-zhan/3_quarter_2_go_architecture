package main

import (
	"fmt"

	"hw5/structures"
)

func main() {
	l := structures.List{}

	fmt.Println("Queue...")
	queue := &structures.Queue{
		List: &l,
	}
	for i := 1; i <= 5; i++ {
		queue.Push(i)
		fmt.Println(i)
	}
	for pop := queue.Pop(); pop != nil; pop = queue.Pop() {
		fmt.Println(pop)
	}

	fmt.Println()
	fmt.Println("Stack...")
	stack := &structures.Stack{
		List: &l,
	}
	for i := 1; i <= 5; i++ {
		stack.Push(i)
		fmt.Println(i)
	}

	for pop := stack.Pop(); pop != nil; pop = stack.Pop() {
		fmt.Println(pop)
	}
}
