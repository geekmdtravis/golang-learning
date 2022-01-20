package ex5

import (
	"fmt"

	"github.com/geekmdtravis/golang-learning/section_5_exercises/ex4"
)

var x ex4.DLen
var y int

func DoExercise() {
	x = 42
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}