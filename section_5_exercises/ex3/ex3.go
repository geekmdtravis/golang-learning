package ex3

import "fmt"

var x int
var y string
var z bool

func DoExercise() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}