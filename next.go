// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 3, 2}
// 	a := validMountainArray(slice)
// 	fmt.Println(a)
// }

// func validMountainArray(arr []int) bool {
// 	if len(arr) < 3 {
// 		return false
// 	}
// 	l := 0
// 	r := len(arr) / 2
// 	fmt.Println(r)
// 	for l < r {
// 		if l < r {
// 			if r != len(arr)/2 {
// 				l++
// 				r++
// 			} else {
// 				break
// 			}

// 		} else {
// 			return false
// 		}
// 	}

// 	l2 := len(arr) / 2
// 	r2 := len(arr[len(arr)/2:])
// 	fmt.Println("IT IS R2 ", r2)

// 	for l2 < r2 {
// 		if l2 > r2 {
// 			if r != len(arr[len(arr)/2:]) {
// 				l2++
// 				r2++
// 			} else {
// 				break
// 			}

// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// package main
// import "fmt"

// func main(){
// 	slice := []int{1,3,5,6,2,1} // 6
// 	a:= check(slice)
// 	fmt.Println(a)
// }

// func c(slice[]int)bool{
// 	l := 0
// 	r := len(slice) - 1

// 	for l < r{
// 		if l < l[l + 1] && r < r[r-1]{
// 			l+=1
// 			r+=1
// 		}else{
// 			return false
// 		}
// 	}
// 	return true
// }

package main

import "fmt"

func main() {
	slice := []int{3, 5, 6, 7, 8, 1}
	a := mountain(slice)
	fmt.Println(a)
}

func mountain(slice []int) bool {
	i := 0
	for i < len(slice)-1 && slice[i] < slice[i+1] {
		i++
	}

	if i == 0 || i == len(slice)-1 {
		return false
	}

	for i < len(slice)-1 && slice[i] > slice[i+1] {
		i++
	}
	return i == len(slice)-1
}
