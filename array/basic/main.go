package main

// MaxUint - maximum unsigned integer value
const MaxUint = ^uint64(0)

// MinUint - minimum unsigned integer value
const MinUint = 0

// MaxInt - maximum integer value
const MaxInt = int(^uint(0) >> 1)

// MinInt - minimum integer value
const MinInt = -MaxInt - 1

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

func minMax(arr []int) []int {
	min := MaxInt
	max := MinInt

	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return []int{min, max}
}

func main() {
	// a := []int{1, 2, 3, 4, -1, 1000000, -45}
	// fmt.Println(reverseArray((a)))
	// fmt.Println(minMax(a))

}
