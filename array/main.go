package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"

	fmt.Println(names)

	// pengisian array saat deklarasi variable
	var fruits = [4]string{"apple", "orange", "banana", "pear"}
	fmt.Println("Panjang Array fruits: ", len(fruits))
	fmt.Println(fruits)

	// inisialisasi array, dengan gaya vertikal
	var numbers = [4]int{
		1,
		2,
		3,
		4,
	}
	fmt.Println(numbers)

	// inisialisasi array dengan length yang bebas
	// [...] menggunakan(.) 3
	var numbers2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Panjang Array numbers2: ", len(numbers2))
	fmt.Println(numbers2)

	// Array Multidimensi
	var numbers3 = [3][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
	}

	fmt.Println(numbers3)
	fmt.Println(numbers3[1][1])

	var numbers4 = [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(numbers4)
	fmt.Println(len(numbers4[0]))

	// Perulangan array menggunakan keyword for
	var fruits1 = [4]string{"apple", "orange", "banana", "pear"}
	for i := 0; i < len(fruits1); i++ {
		fmt.Printf("Elemen %d, %s\n", i, fruits1[i])
	}

	// Perulangan array menggunakan keyword for range
	for i, fruit := range fruits1 {
		fmt.Printf("Elemen[%d], %s\n", i, fruit)
	}

	// Penggunaan variable underscore(_) pada perulangan for range
	for _, fruit1 := range fruits1 {
		fmt.Printf("Nama buah : %s\n", fruit1)
	}

	// Alokasi Elemen Array menggunakan keyword make
	var fruits3 = make([]string, 3)
	fruits3[0] = "Durian"
	fruits3[1] = "Nangka"
	fruits3[2] = "Cerry"

	fmt.Println("Panjang elemen array fruits3 :", len(fruits3))
	fmt.Println(fruits3)

}
