package main

import "fmt"

func main() {
	slice := []int{2, 3, 4, 6}
	slice2 := []int{34, 6, 76, 78}
	n := 1
	m := 1
	a := merge(slice, n, slice2, m)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	a := nums1[:m]
	a2 := nums2[:n]

	res := append(a2, a...)
	

	return res
}
