package main

import "fmt"

func DowhileLoop() {
	number := 3

	// loop that runs infinitely
	for {

		// fmt.Println("Inside the loop")

		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}
}
