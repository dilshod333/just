package main

import "fmt"

func main() {
	slice := []int{1, 1, 2}
	a, a2 := removeDuplicates(slice)
	fmt.Println(a, a2)
}

func removeDuplicates(nums []int) (int, []int) {
	slice := []int{}
	map1 := make(map[int]bool)
	count := 0

	for _, num := range nums {
		if map1[num] {
			continue
		} else {
			map1[num] = true
			slice = append(slice, num)
			count++
		}
	}
	nums = slice
	copy(nums, slice)
	fmt.Println("nums ", nums)
	return count, slice
}
