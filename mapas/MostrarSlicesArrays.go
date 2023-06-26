package mapas

import "fmt"

func MostrarSlices() {
	// one way of creating a slice with the word make
	// it is a function that takes 3 arguments
	// 1. the type of the slice
	// 2. the length of the slice
	// 3. the capacity of the slice
	elements := make([]int, 5, 20)
	// then to add elements to the slice we use the append function
	elements = append(elements, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// another way of creating a slice
	// if we want to speficy the capacity we can add a number in the []
	otherEl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for el, item := range otherEl {
		fmt.Println(el, item)
	}

	// remember that arrays have a fixed length and slices are dynamic
	// fmt.Println(elements)
	// fmt.Println(otherEl)
}
