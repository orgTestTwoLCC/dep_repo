package main_one

import "fmt"

type Greeting struct{}

func (g Greeting) SayHello() {
	fmt.Println("Hello world!")
}
