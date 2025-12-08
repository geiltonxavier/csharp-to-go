package main

import "fmt"

type Box[T any] struct {
	value T
}

func NewBox[T any](v T) Box[T] {
	return Box[T]{value: v}
}

func (b Box[T]) Print() {
	fmt.Printf("boxed value: %v\n", b.value)
}

func PrintAll[T any](items []T) {
	for _, item := range items {
		fmt.Printf("item (%T): %v\n", item, item)
	}
}

func main() {
	textBox := NewBox("hello Go")
	textBox.Print()

	numbers := []int{1, 2, 3}
	PrintAll(numbers)

	words := []string{"go", "csharp"}
	PrintAll(words)
}
