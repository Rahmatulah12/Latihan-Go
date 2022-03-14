package main

import "fmt"

// Pointer adalah reference atau alamat memori dari suatu nilai

/*
	Variabel bertipe pointer ditandai dengan adanya tanda asterisk ( * ) tepat sebelum,
	penulisan tipe data.
	Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai
	yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.

	Ada dua hal penting yang perlu diketahui mengenai pointer:
	Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand ( & )
	tepat sebelum nama variabel.
	Metode ini disebut dengan referencing.
	Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan
	tanda asterisk ( * ) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
*/

func main() {
	numberA := 10               // variabel biasa
	var numberB *int = &numberA // variabel pointer

	fmt.Println("numberA (value) :", numberA)    // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220, mengakses pointer
	fmt.Println("numberB (value) :", *numberB)   // 4, mengakses nilai nya
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	numberA = 5
	fmt.Println("numberA (value) :", numberA)    // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220, mengakses pointer
	fmt.Println("numberB (value) :", *numberB)   // 4, mengakses nilai nya
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	changeValue(numberB, 20)
	fmt.Println("numberB (value) :", *numberB) // 4
}

func changeValue(number *int, newValue int) {
	*number = newValue
}
