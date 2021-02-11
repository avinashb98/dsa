/*
Move all negative numbers to beginning and positive to end with constant extra space

An array contains both positive and negative numbers in random order. Rearrange the array elements so that all negative numbers appear before all positive numbers.
Examples :

Input: -12, 11, -13, -5, 6, -7, 5, -3, -6
Output: -12 -13 -5 -7 -3 -6 11 6 5
*/

package main

import "fmt"

func rearrange(arr []int) {
	j := 0
	n := len(arr)
	var i int
	for i = 0; i < n; i++ {
		if arr[i] < 0 {
			if i != j {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
			j++
		}
	}
}

func moveNegativeToLeft(arr []int) {
	left := 0
	right := len(arr) - 1
	var tmp int
	for left < right {
		if arr[right] < 0 {
			for arr[left] < 0 && left < right {
				left++
			}
			if !(left < right) {
				break
			}
			tmp = arr[right]
			arr[right] = arr[left]
			arr[left] = tmp
			left++
		}
		right--
	}
}

func main() {
	arr := []int{-12, 11, -13, -5, 6, -7, 5, -3, -6}
	// moveNegativeToLeft(arr)
	rearrange(arr)
	fmt.Println(arr)
}
