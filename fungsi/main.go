package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	printMessage("Hello", []string{"Rahmatulah", "Sidik"})

	// generate angka acak dengan range
	var randomValue int
	randomValue = randomNumberWithRange(10, 1)
	fmt.Println("Random number: ", randomValue)

	// penggunaan keyword return untuk memberhentikan proses
	divideNumber(10, 0)
	divideNumber(4, -2)
	fmt.Println(decisionNumber(0))
	fmt.Println(divisibleBy2(4, 33))

	// multi return
	luas, keliling := calculateAreaAndRoundOfSquare(5)
	fmt.Println("Luas: ", luas, " Keliling: ", keliling)

	fmt.Println(returnCombineSliceAndStruct()...)

	fmt.Println(predefinedReturnValue("Cok"))

	fmt.Println("Luas Lingkaran: ", calcCircleArea(5))
	fmt.Println("Keliling Lingkaran: ", calcCircleRound(5))

	// variadic function
	fmt.Println("Total harga: ", SumPrices(100, 200, 300, 400))

	// Casting pada Go
	fmt.Println(float64(8))

	// Pengisian parameter variadic function with slice
	var prices1 = []uint16{1000, 2000, 3000, 4000}
	fmt.Println("Total harga dari parameter slice :", SumPrices(prices1...))

	hobbies("Rahmatulah", "Membaca", "Ngoding", "Olahraga")
	var hobbies1 = []string{"Membaca", "Ngoding", "Olahraga"}
	hobbies("Rahmatulah", hobbies1...)

	// Closure Function
	var getMinMax = func(numbers []int) (min, max int) {

		for i, value := range numbers {
			switch {
			case i == 0:
				min, max = value, value
			case value > max:
				max = value
			case value < min:
				min = value
			}
		}
		return
	}

	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := getMinMax(numbers)

	// Penggunaan Template String %v
	// Template %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya.
	fmt.Printf("Min: %v, Max: %v\n", min, max)

	// Immediately-Invoked Function Expression (IIFE)
	/*
		Closure jenis ini dieksekusi langsung pada saat deklarasinya.
		Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali,
		bisa mengembalikan nilai, bisa juga tidak.
	*/
	// func membuang angka yang lebih kecil dari parameter
	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// membuang angka yang lebih kecil dari 3
	var newNumbers2 = func(min int) []int {
		var temp []int
		for _, value := range numbers2 {
			if value >= min {
				temp = append(temp, value)
			}
		}
		return temp
	}(3) // ciri khas IIFE adanya tanda kurung parameter setelah deklarasi function

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers2)

}

func printMessage(message string, array []string) {
	fmt.Println(message, strings.Join(array, " "))
}

func randomNumberWithRange(min, max int) int {
	var value int
	value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(a, b int) {
	if b == 0 {
		fmt.Println("Error: angka pembagi tidak boleh 0")
		return
	}

	res := a / b
	fmt.Println("Hasil bagi: ", res)
}

func decisionNumber(a int) int {
	if a == 0 {
		return 0
	}
	return a
}

func divisibleBy2(a, b uint8) bool {
	if a%b == 0 {
		return true
	}
	return false
}

// calculate area and round of square
func calculateAreaAndRoundOfSquare(a uint8) (uint8, uint8) {
	var area, round uint8
	area = a * a
	round = a + a
	return area, round
}

func returnMultipleStringWithMap() map[string]string {
	var mapString = make(map[string]string)
	mapString["nama"] = "Rahmatulah"
	mapString["alamat"] = "Jakarta"
	return mapString
}

func returnDataSliceInt() []int {
	var dataSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return dataSlice
}

func returnDataStruct() []struct {
	nama   string
	alamat string
} {
	var dataStruct = []struct {
		nama   string
		alamat string
	}{
		{"Rahmatulah", "Jakarta"},
		{"Sidik", "Bandung"},
	}
	return dataStruct
}

func returnStructWithParamateStruct(data struct {
	nama   string
	alamat string
}) struct {
	nama   string
	alamat string
} {
	return data
}

func returnCombineSliceAndStruct() []interface{} {
	var dataSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var dataStruct = []struct {
		nama   string
		alamat string
	}{
		{"Rahmatulah", "Jakarta"},
		{"Sidik", "Bandung"},
	}
	var dataCombine = []interface{}{dataSlice, dataStruct}
	return dataCombine
}

func predefinedReturnValue(inputName string) (name string) {

	if inputName != "Rahmatulah" {
		return "Nama tidak ditemukan"
	}
	return
}

func calcCircleArea(r float64) float64 {
	pi := math.Pi
	return pi * r * r
}

func calcCircleRound(r float64) float64 {
	pi := math.Pi
	return pi * 3.14 * r
}

func SumTotalFromSlice(data []int) int {
	var total int
	for _, value := range data {
		total += value
	}
	return total
}

// Variadic function, fungsi dengan parameter tak terbatas
func SumPrices(prices ...uint16) uint16 {
	var total uint16
	for _, value := range prices {
		total += value
	}

	return total
}

// fungsi dengan kombinasi parameter variadic dan biasa
func hobbies(name string, hobbies ...string) {
	var joinHobbiesToString string = strings.Join(hobbies, ",\n")

	fmt.Println("Nama: ", name)
	fmt.Println("Hobi: ", joinHobbiesToString)

}
