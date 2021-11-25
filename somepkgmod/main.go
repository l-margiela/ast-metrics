package main

import (
	"fmt"
	"time"
)

func Greet(name string, age int) error {
	metricsStart4 := time.Now()
	fmt.Println("Let me evaluate the input…")
	if name == "" {
		metricsStart1 := time.Now()
		time.Sleep(time.Second)
		fmt.Println("Empty name")
		fmt.Printf("Block time measurement (ID 1); time: %s\n", time.Since(metricsStart1))
	}

	if age < 1 {
		metricsStart2 := time.Now()
		time.Sleep(time.Second * 2)
		fmt.Println("Invalid age")
		fmt.Printf("Block time measurement (ID 2); time: %s\n", time.Since(metricsStart2))
	}

	switch {
	case age < 13:
		metricsStart7 := time.Now()
		fmt.Printf("Hi, %s. You're a kid.\n", name)
		fmt.Printf("Block time measurement (ID 7); time: %s\n", time.Since(metricsStart7))
	case age < 18:
		metricsStart8 := time.Now()
		fmt.Printf("Hello, %s. You're a teenager.\n", name)
		fmt.Printf("Block time measurement (ID 8); time: %s\n", time.Since(metricsStart8))
	default:
		metricsStart9 := time.Now()
		fmt.Println("Thinking…")
		time.Sleep(time.Second * 3)
		fmt.Printf("Howdy, %s. You're an adult.\n", name)
		fmt.Printf("Block time measurement (ID 9); time: %s\n", time.Since(metricsStart9))
	}
	fmt.Printf("Block time measurement (ID 4); time: %s\n", time.Since(metricsStart4))

	return nil
}

func main() {
	metricsStart6 := time.Now()
	if err := Greet("Adam", 18); err != nil {
		metricsStart5 := time.Now()
		fmt.Println(err)
		fmt.Printf("Block time measurement (ID 5); time: %s\n", time.Since(metricsStart5))
	}
	fmt.Printf("Block time measurement (ID 6); time: %s\n", time.Since(metricsStart6))
}
