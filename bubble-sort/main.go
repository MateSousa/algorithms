package main

import "fmt"

func main() {
	BubbleSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	BubbleSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	BubbleSort([]int{1, 33, 19, 20, 4, 3, 0, -129, 203, 2, 1, 4, 5, 6, 7, 8, 9})
	BubbleSort([]int{})
}

func BubbleSort(arr []int) []int {
	r := arr
	if len(r) == 0 {
		return []int{}
	}

	if len(r) == 1 {
		return r
	}
	// doesnt work :(
	// for i := range r {
	//
	// if i == len(r)-2 {
	// if r[i] > r[i+1] {
	// r[i], r[i+1] = r[i+1], r[i]
	// }
	// break
	// }
	//
	// if r[i] < r[i+1] {
	// continue
	// }
	//
	// if r[i] == r[i+1] {
	// continue
	// }
	//
	// if r[i] > r[i+1] {
	// r[i], r[i+1] = r[i+1], r[i]
	// }
	//
	// }
	fmt.Println(r)
	return r
}
