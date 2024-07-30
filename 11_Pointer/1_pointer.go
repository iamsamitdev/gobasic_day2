package main

import "fmt"

func main() {

	// สร้างตัวแปร 2 ตัว
	i, j := 42, 2701

	// สร้างตัวแปร pointer ชี้ไปที่ i
	p := &i // address of i like this: 0xc0000160e0

	// อ่านค่า i ผ่าน pointer
	fmt.Println(*p) // 42

	// เปลี่ยนค่า i ผ่าน pointer
	*p = 21
	fmt.Println(i) // 21

	// สร้างตัวแปร pointer ชี้ไปที่ j
	p = &j // address of j like this: 0xc0000160e8

	// เปลี่ยนค่า j ผ่าน pointer
	*p = *p / 37
	fmt.Println(j) // 73
}
