package ex3

import "fmt"

func DoExercise() {
	fv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(fv[2:4])
	fmt.Println(fv[4:6])
	fmt.Println(fv[1:9])
	fmt.Println(fv[:9])
	
}