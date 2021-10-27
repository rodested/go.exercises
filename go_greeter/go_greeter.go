package main

import (
	"github.com/rodested/golib_greet"
	"fmt"
)

func main() {
	//Get a greeting
	message := golib_greet.Hello("Horsch")
	fmt.Println(message)
}
