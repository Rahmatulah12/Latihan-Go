package main

import "fmt"

func main() {
	/**
	fmt.Println("Hello World")
	fmt.Println("Hello", "Rahmatulah", "Sidik")
	*/

	// Variable
	/*
		var firstName string = "Rahmatulah"
		lastName := "Sidik"
		var coba = true
		var age int
		age = 27

		fmt.Printf("Hello %s %s. Umur saya %d tahun.", firstName, lastName, age)
		fmt.Println("Hello", firstName, lastName, "Umur saya", age, "tahun.")
		fmt.Println(coba)
	*/

	/*
		var first, second, third int = 1, 2, 3
		fmt.Println(first, second, third)

		//  Variable _ (undescore), adalah variable temporary yang tidak digunakan
		var firstName, lastName, _ = "Rahmatulah", "Sidik", "Coba"
		fmt.Println(firstName, lastName)
	*/

	//  Deklarasi Variable Menggunakan keyword new (untuk membuat pointer)
	/*
		name := new(string)

		fmt.Println(name) // 0xc000010230 => alamat memori
		fmt.Println(*name) // isi variable menggunakan tanda asterik(*)
	*/

	/*
		Deklarasi varibale menggunakan keyword make
		Keyword ini hanya bisa digunakan untuk :
			- channel
			- slice
			- map
	*/

	/*
		Tipe Data Numerik Non-Desimal
			- uint => bilangan positif
			- int => bilangan positif & negatif
	*/
	/*
		var postiveNumber = 10
		negativeNumber := -10

		fmt.Printf("Bilangan Positif : %d\n", postiveNumber)
		fmt.Printf("Bilangan Negatif : %d\n", negativeNumber)
	*/

	/*
		Tipe Data Numerik Desimal
			- float64
			- float32
	*/
	/*
		var decimalNumber float32 = 3.14
		fmt.Printf("Bilangan Desimal : %f\n", decimalNumber)
		fmt.Printf("Bilangan Desimal : %.3f\n", decimalNumber)
	*/

	/*
		Tipe Data Boolean
			- true
			- false
	*/
	/*
		exists := true
		if !exists {
			fmt.Printf("exists ? %t\n", exists)
		} else {
			fmt.Printf("exists ? %t\n", exists)
		}

		var words string = `Hallo, nama saya "Rahmatulah Sidik".
			Malam ini adalah malam Jum'at.
		`

		fmt.Println(words)
	*/

	/*
		Nilai nill & zero value
	*/

	/*
		// Konstanta
		const FULLNAME string = "Rahmatulah Sidik"
		fmt.Println(FULLNAME)

		// Operator
		var value = (((2+6)%3)*4 - 2) / 3
		fmt.Print(value, "\n")
	*/

	// Condition
	var nilai = 6
	// condition when nilai is equal 6, grade is C
	if nilai == 6 {
		fmt.Println("Grade : C")
	}

	// variable temporary pada if-else
	var point int8 = 70
	if percent := point / 10; percent >= 10 {
		fmt.Printf("%d, nilai sempurna.\n", percent)
	} else if percent >= 70 {
		fmt.Printf("%d, nilai cukup.\n", percent)
	} else {
		fmt.Printf("%d, nilai kurang.\n", percent)
	}

	// switch case
	switch point {
	case 100:
		fmt.Println("Grade A+")
	case 90:
		fmt.Println("Grade A")
	case 80:
		fmt.Println("Grade B")
	case 70:
		fmt.Println("Grade C")
	case 60:
		fmt.Println("Grade D")
	default:
		fmt.Println("Grade E")
	}

	// case banyak kondisi
	switch point {
	case 100, 90:
		fmt.Println("Grade A")
	case 80, 70:
		fmt.Println("Grade B")
	case 60, 50:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade E")
	}

	switch point {
	case 100, 90:
		fmt.Println("Grade A")
	case 80, 70, 65:
		fmt.Println("Grade B")
	default:
		{
			if point <= 60 {
				fmt.Println("Grade C")
			} else {
				fmt.Println("Grade E")
			}
		}
	}

	// switch case dengan gaya if-else
	switch {
	case point == 100:
		fmt.Println("Grade A+")
	case (point <= 90) && (point >= 80):
		fmt.Println("Grade A")
	case (point <= 80) && (point >= 70):
		fmt.Println("Grade B")
	case (point <= 70) && (point >= 60):
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade E")
	}

	// penggunaan keyword fallthrough
	// fallthrough akan melewati kondisi selanjutnya
	switch {
	case point == 100:
		fmt.Println("Grade A+")
	case (point <= 90) && (point >= 80):
		fmt.Println("Grade A")
	case (point <= 80) && (point >= 70):
		fmt.Println("Grade B")
	case (point <= 70) && (point >= 60):
		fmt.Println("Grade C")
		fallthrough
	default:
		fmt.Println("Grade E")
	}

	// kondisi bersarang
	if point > 70 {
		switch point {
		case 80:
			fmt.Println("Grade B+")
		case 90:
			fmt.Println("Grade A")
		default:
			fmt.Println("Grade A+")
		}
	} else {
		if point <= 70 {
			fmt.Println("Grade B")
		} else if point <= 60 {
			fmt.Println("Grade C")
		} else {
			fmt.Println("Grade E")
		}
	}
}
