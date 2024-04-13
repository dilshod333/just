package main

import "fmt"

func main() {
	slice := []int{12, 344, 45, 5676}
	count := 0
	for i := 0; i < len(slice); i++ {
		if check(slice[i])%2 == 0 {
			count++
		}
	}
	fmt.Println(count)

}

func check(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count
}
