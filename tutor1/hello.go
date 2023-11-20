package main

import (
	"fmt"
)

// ประกาศตัวแปร x แบบ global และกำหนดค่าเป็น 5
var x int = 5

// ประกาศตัวแปร y แบบ global โดยไม่กำหนดค่า (ค่าเริ่มต้นคือ 0 สำหรับ int)

// ประกาศฟังก์ชันชื่อ Inti (ที่มีความน่าจะเป็นเป็นการพิมพ์ผิด, น่าจะหมายถึง "Init")
func Inti() {
	// พิมพ์ค่าของตัวแปร global x
	fmt.Println("x =", x)
}

// ฟังก์ชัน main, ที่เป็นจุดเริ่มต้นของการทำงานของโปรแกรม
func main() {

	// ประกาศตัวแปร boolean ชื่อ t โดยไม่กำหนดค่า (ค่าเริ่มต้นคือ false)
	var t bool

	// ประกาศและกำหนดค่าตัวแปร boolean ชื่อ f เป็น true
	var f bool = true

	// พิมพ์ค่าของ t และ f
	fmt.Println("t is", t)
	fmt.Println("f is", f)

	// ประกาศและกำหนดค่าตัวแปร integer ชื่อ age เป็น 21
	var age int = 21

	// ประกาศและกำหนดค่าตัวแปร float-point ชื่อ favNum เป็น 1.6180339
	var favNum float64 = 1.6180339

	// ประกาศและกำหนดค่าตัวแปร complex number ชื่อ complex เป็น 5 + 5i
	var complex complex128 = 5 + 5i

	// ประกาศและกำหนดค่าตัวแปร rune (alias for int32) ชื่อ r เป็น 10
	var r rune = 10

	// พิมพ์ค่าของ age, favNum, complex, และ r
	fmt.Println("age is", age, "favNum is", favNum)
	fmt.Println("complex128 is", complex)
	fmt.Println("rune is", r)

	// ประกาศตัวแปร string ชื่อ myName และกำหนดค่าเป็น "kanya"
	var myName string = "kanya"

	// พิมพ์ string concatenation โดยใช้ myName
	fmt.Println(myName + " is a robot")

	// พิมพ์ตัวอักษรที่ตำแหน่งที่ 3 ของ myName
	fmt.Println(myName[3])

	// พิมพ์ความยาวของ myName
	fmt.Println("length of myName is", len(myName))

	// ประกาศ array ชื่อ arry5 ที่มีขนาด 5 และมี elements เป็น float64
	var arry5 [5]float64

	// กำหนดค่าแต่ละ elements ของ arry5
	arry5[0] = 98.7
	arry5[1] = 93.2
	arry5[2] = 77.2
	arry5[3] = 83.7
	arry5[4] = 88.2

	// พิมพ์ arry5 ทั้งหมด
	fmt.Println(arry5)

	// พิมพ์ความยาวของ arry5
	fmt.Println("length of arry is", len(arry5))

	// พิมพ์ค่าของ arry5 ที่ index 3
	fmt.Println("arry[3] is ", arry5[3])

	// ประกาศและกำหนดค่า array ชื่อ arry3 ที่มีขนาด 3 และมี elements เป็น float64
	arry3 := [3]float64{98, 93, 77}

	// พิมพ์ arry3 ทั้งหมด
	fmt.Println(arry3)

	// ประกาศ slice ชื่อ s ซึ่งเป็นการอ้างอิงถึงส่วนหนึ่งของ arry5 (index 0 ถึง 1)
	var s []float64 = arry5[0:2]

	// พิมพ์ slice s
	fmt.Println(s)