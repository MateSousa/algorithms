package main

import (
	"fmt"
	"log"
)

func main() {
	Fibonacci(10)
	Fibonacci(3)
	Fibonacci(0)
	Fibonacci(-5)
}

func Fibonacci(n int) ([]int, error) {
	r := []int{}
	prev, next := 0, 1

	if n < 0 {
		log.Fatal("n must be a positive number")
	}
	if n == 0 {
		fmt.Println(r)
		return r, nil
	}

	for i := 0; i < n; i++ {
		r = append(r, prev)
		prev, next = next, prev+next
	}
	fmt.Println(r)
	return r, nil
}
