package ex5

import "fmt"

func DoExercise() {
	fv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fv = append(fv[:5], fv[7:]...)

	fmt.Println(fv)
}