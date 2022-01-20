package ex1

import "fmt"

func DoExercise() {
	var fv [5]int = [5]int{1, 2, 3, 4, 5}

	for i, v := range fv {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", fv)
}