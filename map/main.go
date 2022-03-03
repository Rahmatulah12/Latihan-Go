package main

import "fmt"

func main() {
	months := map[string]int{} // default nilai awal map adalah nil

	months["January"] = 31
	months["February"] = 28
	months["March"] = 31

	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println("Jumlah hari pada bulan Januari adalah", months["January"])
	fmt.Println("Jumlah hari pada bulan Februari adalah", months["February"])
	fmt.Println("Jumlah hari pada bulan March adalah", months["March"])
	fmt.Println("Jumlah hari pada bulan April adalah", months["April"]) // nilai default int = 0, karena key april blm diset

	//  Inisialisasi Awal Map
	months2 := map[string]int{"januari": 31, "februari": 28}
	months3 := map[string]int{
		"januari":  31,
		"februari": 28,
	}

	fmt.Println(months2)
	fmt.Println(months3)

	/*
		Variabel map bisa di-inisialisasi dengan tanpa nilai awal, caranya menggunakan tanda kurung kurawal, contoh: map[string]int{} .
		Atau bisa juga dengan menggunakan keyword make dan new . Contohnya bisa dilihat pada kode berikut. Ketiga cara di bawah ini
		intinya adalah sama.
	*/
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// Iterasi Item Map menggunakan keyword for-range
	for key, value := range months {
		fmt.Println(key, "\t: ", value)
	}

	fmt.Println("Variable Map months2")
	for key, value := range months2 {
		fmt.Println(key, "\t: ", value)
	}

	//  Menghapus item pada map, delete()
	delete(months, "February")
	delete(months2, "januari")

	fmt.Println(months)  // map[January:31 March:31]
	fmt.Println(months2) // map[februari:28]

	// cek keberadaan key pada Map
	var value, exists = months3["januari"]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key januari tidak ditemukan")
	}

	// Kombinasi Slice & Map
	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }

	// Penyederhanaan penulisan
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
