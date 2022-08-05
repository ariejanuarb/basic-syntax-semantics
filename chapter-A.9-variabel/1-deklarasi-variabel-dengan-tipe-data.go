package main

import "fmt"

/*
3 komponen utama variabel dalam golang :
1. type -> jenis informasi yg dsimpan dalam variabel
2. value -> value yg disimpan
3. address -> dimana variabel bisa ditemukan pada memori komputer

why we need variable ?
*/

func main() {
	// manifest typing = ketika deklrasi, tipe data harus ditulis juga
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
}
