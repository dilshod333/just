// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "hello salom  nima ,. qaley alo qaley"
// 	a := strings.Split(str, " ")
// 	map1 := make(map[string]int)
// 	check := ".,%$"
// 	for i := 0; i < len(a); i++ {

// 		if strings.Contains(a[i], check) {
// 			continue
// 		}else{
// 			map1[a[i]]++
// 		}
// 	}
// 	fmt.Println(map1)
// }


package main 
import "fmt"

type  Person struct{
	first_name 	string,
	last_name 	string,
	age  		int,
	gender 		string,
	phone_number string ,
	address 	string, 
}

func main(){
	residence := []Person{}

	tesha := Person{

	}
}