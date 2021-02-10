/*
Sort an array of 0s, 1s and 2s

Given an array of size N containing only 0s, 1s, and 2s; sort the array in ascending order.


Example 1:

Input:
N = 5
arr[]= {0 2 1 2 0}
Output:
0 0 1 2 2
Explanation:
0s 1s and 2s are segregated
into ascending order.
Example 2:

Input:
N = 3
arr[] = {0 1 0}
Output:
0 0 1
Explanation:
0s 1s and 2s are segregated
into ascending order.

You don't need to read input or print anything. Your task is to complete the function sort012() that takes an array arr and N as input parameters and sorts the array in-place.

Expected Time Complexity: O(N)
Expected Auxiliary Space: O(1)


Constraints:
1 <= N <= 10^5
0 <= A[i] <= 2

*/

package main

import "fmt"

func sort012(arr []int) []int {
	count0 := 0
	count1 := 0
	count2 := 0

	for _, val := range arr {
		if val == 0 {
			count0++
		} else if val == 1 {
			count1++
		} else {
			count2++
		}
	}

	ptr := 0
	var correctVal int

	for count0 > 0 || count1 > 0 || count2 > 0 {
		if count0 > 0 {
			correctVal = 0
			count0--
		} else if count1 > 0 {
			correctVal = 1
			count1--
		} else {
			correctVal = 2
			count2--

		}
		arr[ptr] = correctVal
		ptr++
	}

	return arr
}

func main() {
	arr := []int{1, 0, 0, 0, 2, 2, 1, 2, 1, 1, 0}
	fmt.Println(sort012(arr))
}
