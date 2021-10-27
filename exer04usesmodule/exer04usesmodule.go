package main

import (
	"exercises/exer03amodule"
	"fmt"
)

func main() {
	//Get a greeting
	message := exer03amodule.Hello("Horsch")
	fmt.Println(message)
}
