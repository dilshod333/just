package main

import "fmt"

func main() {
	arr := []int{1,0,2,3,0,4,0}
	slice := []int{}
	fmt.Println(slice)
	for i := 0; i < len(arr); i++ {
		if len(slice) < len(arr) {
			if arr[i] != 0 {
				slice = append(slice, arr[i])
			} else if arr[i] == 0 {
				slice = append(slice, 0)
				slice = append(slice, 0)
			}
		}
	}
	for i := 0; i < len(slice); i++ {
		arr[i] = slice[i]
	}
	fmt.Println(slice, arr)

}
