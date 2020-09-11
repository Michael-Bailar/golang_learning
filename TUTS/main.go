package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey")
	random()
}

func random() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The function recovered from panic with a reason: %s", r)
		}
	}()

	fmt.Println("1")
	panic("Some Reason")
	fmt.Println("2")
}
