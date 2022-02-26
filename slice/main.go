package main

import "fmt"

func main() {
	// inisialisasi slice
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	// Perbedaan slice dan array
	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)

	var fruitsD = [...]string{"apple", "grape", "banana", "melon", "papaya", "nangka"}
	// membuat slice dari array
	var sliceFruitsD = fruitsD[1:4]
	var sliceFruitsE = fruitsD[:4]
	fmt.Println(sliceFruitsD)

	/*
	* Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice
	* yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan
	* berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
	 */
	sliceFruitsD[0] = "mangga"
	fmt.Println(sliceFruitsD)
	fmt.Println(sliceFruitsE)
}
