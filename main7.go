package main

import "fmt"

type animal struct {
	id   string
	name string
	age  int
}

func main() {

	// slice2 := []int{6}
	// fmt.Println(cap(slice2), slice2)
	s := []int{}
	s1 := make([]int, 0)
	fmt.Println(s == nil, s1 == nil, 0 == 0)
	slice3 := make([]int, 4)
	slice3 = append(slice3, 5, 8, 11, 22)
	slice3 = append(slice3, 44, 50, 77, 44, 55, 66)
	slice3 = append(slice3, 45, 99, 0)
	y := []int{}
	copy(y, slice3)
	fmt.Println(slice3, cap(slice3), len(slice3))

	fmt.Println(y)
	var Animal animal = animal{
		name: "tiger",
		age:  18,
		// id:   "30200",
	}
	fmt.Println(Animal)

	Animal.name = "geopard"
	fmt.Println(Animal.name, Animal.age, Animal.id)

}
