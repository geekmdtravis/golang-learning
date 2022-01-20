package ex4

import "fmt"

func DoExercise() {
	xa := []int{1, 2, 3}
	xb := []int{7, 8, 9}
	
	xc:= append(xa, xb...)
	fmt.Println(xc)

	xd := append(xc, 5, 6, 7)
	fmt.Println(xd)
}