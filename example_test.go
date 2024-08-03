package greetings_test

import (
	"fmt"

	"github.com/seaboie/greetings"
)

func Example() {
	name := "Kritbovorn"
	message := greetings.Hello(name)
	fmt.Println(message)
	// Output:
	// Hi nice to meet you, Kritbovorn. and welcome to thailand.
}