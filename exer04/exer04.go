package main

import (
	"fmt"

	"github.com/rodested/golibgreet"
)


func main() {
	//Get a greeting
	message := golibgreet.Hello("Horsch")
	fmt.Println(message)
}
