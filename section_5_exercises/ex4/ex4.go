package ex4

import "fmt"

type DLen int

func DoExercise() {
	var x DLen 
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)

}