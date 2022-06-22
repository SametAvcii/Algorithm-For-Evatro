package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{3, 4, 9, 14, 15, 19, 28, 37, 47, 50, 54, 56, 59, 61, 70, 73, 78, 81, 92, 95, 97, 99}
	fmt.Println(FindSubSet(arr))
}

func findBigger(n []int) int { //Find the bigger number of input array
	lenght := len(n)
	big := 0
	for i := 0; i < lenght; i++ {
		if n[i] > big {
			big = n[i]
		}
	}
	return big
}

func FindSubSet(arr []int) int { //find subset in the input array
	big := findBigger(arr)
	sum := 0
	count := 0
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		count++
		if sum == big {
			count += 1
			return count
		}
	}
	return FindSubSet(arr[:len(arr)-1])

}
