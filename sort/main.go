package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByPersonName []Person

func (ppl ByPersonName) Len() int {
	return len(ppl)
}

func (ppl ByPersonName) Swap(i int, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func (ppl ByPersonName) Less(i int, j int) bool {
	return ppl[i].Name < ppl[j].Name
}

type ByPersonAge []Person

func (ppl ByPersonAge) Len() int {
	return len(ppl)
}

func (ppl ByPersonAge) Swap(i int, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func (ppl ByPersonAge) Less(i int, j int) bool {
	return ppl[i].Age < ppl[j].Age
}


func main() {
	is := []int{4,3,99,11,3}
	sort.Ints(is)
	fmt.Println(is)

	ps := []Person{{"Bob",33}, {"Alice",44}, {"Fred",55}}
	sort.Sort(ByPersonName(ps))
	fmt.Println(ps)

	sort.Sort(ByPersonAge(ps))
	fmt.Println(ps)
}