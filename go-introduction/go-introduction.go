package main

import "fmt"

func main() {
	fmt.Println("Square, Cube, or SquareCube!")
	fmt.Println("Please input total number to check in integer: ")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= num; i++ {
		square, cube := false, false
		for j := 0; j <= i; j++ {
			if j*j == i {
				square = true
			}
			if j*j*j == i {
				cube = true
			}
		}

		squareCube := square && cube

		switch {
		case squareCube:
			fmt.Println("SquareCube")
		case cube:
			fmt.Println("Cube")
		case square:
			fmt.Println("Square")
		default:
			fmt.Println(i)
		}
	}
}
