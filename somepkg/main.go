package main

import (
	"fmt"
	"time"
)

func Greet(name string, age int) error {
	fmt.Println("Let me evaluate the input…")
	if name == "" {
		time.Sleep(time.Second)
		fmt.Println("Empty name")
	}

	if age < 1 {
		time.Sleep(time.Second * 2)
		fmt.Println("Invalid age")
	}

	switch {
	case age < 13:
		fmt.Printf("Hi, %s. You're a kid.\n", name)
	case age < 18:
		fmt.Printf("Hello, %s. You're a teenager.\n", name)
	default:
		fmt.Println("Thinking…")
		time.Sleep(time.Second * 3)
		fmt.Printf("Howdy, %s. You're an adult.\n", name)
	}

	return nil
}

func main() {
	if err := Greet("Adam", 18); err != nil {
		fmt.Println(err)
	}
}
