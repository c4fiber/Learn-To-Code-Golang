package main

import "fmt"

func composite_literal_slice_literal() {
	x := []int{2, 4, 6,}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 10)
	y = append(y, 777)

	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	// map
	// struct
}