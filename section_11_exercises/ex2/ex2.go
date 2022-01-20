package ex2

import "fmt"

func DoExercise() {
	fv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range fv {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", fv)
}