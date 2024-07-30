package main

import "fmt"

func Forloop() {

	// for loop terminates when i becomes 6
	for i := 1; i <= 10; i++ {
		if i == 5 {
			// reak // หยุดการทำงานของ loop ทันที
			continue // ข้ามการทำงานของ loop ทันที
		}
		fmt.Println(i)
	}

}
