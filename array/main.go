package main

import "fmt"

func reverseArray(arr []int) []int {
	arrayLength := len(arr)
	var i int
	var tmp int
	for i = 0; i < arrayLength/2; i++ {
		tmp = arr[arrayLength-1-i]
		arr[arrayLength-1-i] = arr[i]
		arr[i] = tmp
	}
	return arr
}

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(reverseArray((a)))
}
