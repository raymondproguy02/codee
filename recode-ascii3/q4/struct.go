package main

import "fmt"

type me struct {
	name   string
	age    int
	gender string
}
type you struct {
	name   string
	age    int
	gender string
}

func main() {
	a := me{name: "ray", age: 897, gender: "creator"}
	z := me{name: "cel", age: 1897, gender: "coder"}
	fmt.Println(a.gender, z.age)
}
