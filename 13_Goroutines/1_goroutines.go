package main

import (
	"fmt"
	"time"
)

// สร้างฟังก์ชันตัวอย่างไว้รันใน goroutine
func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Microsecond) // หน่วงเวลา 100 ไมโครวินาที
	}
}

func main() {

	// เรียกใช้ฟังก์ชัน printNumbers ใน goroutine
	go printNumbers()

	// ฟังก์ชันหลักใน Goroutine จะทำงานพร้อมกับฟังก์ชันหลัก
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
		time.Sleep(150 * time.Microsecond) // หน่วงเวลา 150 ไมโครวินาที
	}

}
