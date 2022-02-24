package main

import "fmt"

func main() {
	fmt.Println("Perulangan For")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Perulangan ke-", i)
	// }

	// var angka uint8 = 7
	// // for dengan kondisi
	// for angka > 0 {
	// 	fmt.Println("Perulangan ke-", angka)
	// 	angka--
	// 	if angka == 0 {
	// 		break
	// 	}
	// }

	// for tanpa kondisi, dihentikan dengan keyword break
	// var i = 0
	// for {
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke-", i)
	// }

	// Mencetak angka genap
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke-", i)
	// }

	// // mencetak angka ganjil
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke-", i)
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := i; j <= 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// Pemanfaatan Label Dalam Perulangan
	// perulangan memunculkan matriks
outerLoops:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 4 {
				break outerLoops
			}
			fmt.Println("Matriks[", i, "][", j, "]")
		}
	}
}
