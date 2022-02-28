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

	// fungsi len() => menghitung jumlah elemen slice yang ada
	//  fungsi cap() => menghitung lebar atau kapasitas maksimum slice
	var fruitsE = []string{"pear", "semangka", "buah naga"}
	fmt.Println(len(fruitsE)) // 3
	fmt.Println(cap(fruitsE)) // 3

	var fruitsE2 = fruitsE[0:2]
	fmt.Println(len(fruitsE2)) // 2
	fmt.Println(cap(fruitsE2)) // 3 => karena mengambil dari index ke 0 dari referensi slice sblmnya

	var fruitsE3 = fruitsE[1:3]
	fmt.Println(len(fruitsE3)) // 2
	fmt.Println(cap(fruitsE3)) //2 => karena mengambil dari index ke 1 dari referensi slice sblmnya

	// append(), digunakan untuk menambahkan elemen pada slice
	fruitsF := append(fruitsE, "apel")
	fmt.Println(fruitsF)

	// copy() => mengcopy elemen slice pada element0 (parameter ke 2),  ke element1 (parameter ke 1)
	/*
		Yang ter-copy hanya 3 buah (meski element0 memiliki 4 elements) hal ini karena copy() hanya meng-copy elements sebanyak
		len(element1) .
	*/
	element0 := make([]string, 3)
	element1 := []string{"mangga", "apel", "nangka", "pinaple"}
	newElement := copy(element0, element1)

	fmt.Println(element0)
	fmt.Println(element1)
	fmt.Println(newElement)

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// Pengaksesan Elemen Slice dengan 3 index
	/*
		3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya. Cara menggunakannnya yaitu dengan menyisipkan
		angka kapasitas di belakang, seperti fruits[0:1:1] . Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di
		slicing.
	*/

	var fruitsG = []string{"apple", "grape", "banana", "melon", "papaya", "nangka"}
	var fruitsG2 = fruitsG[0:3]
	var fruitsG3 = fruitsG[0:3:3] // dengan kapasitas 3

	fmt.Println(fruitsG)
	fmt.Println(len(fruitsG))
	fmt.Println(cap(fruitsG))

	fmt.Println(fruitsG2)
	fmt.Println(len(fruitsG2))
	fmt.Println(cap(fruitsG2))

	fmt.Println(fruitsG3)
	fmt.Println(len(fruitsG3))
	fmt.Println(cap(fruitsG3))

}
