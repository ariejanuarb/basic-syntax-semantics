package main

import "fmt"

func main() {
	// 1. deklarasi variabel chicken = map, string = key-nya, int = value-nya
	var chicken map[string]int
	chicken = map[string]int{}

	// 2. deklarasi map menggunakan make function
	var customer = make(map[int]string)
	fmt.Println(customer) // returning "map[]" with empty value

	customer[1] = "owner"
	customer[2] = "member"
	fmt.Println(customer)

	// adding new value and overwriting the old one
	customer[2] = "employee"
	customer[3] = "supervisor"
	fmt.Println(customer)

	// pengisian nilai pada key map, sifatna overwrite
	chicken["januari"] = 50 // key = januari dimasukkan value = 50 (int)
	chicken["februari"] = 40
	chicken["mei"] = 99

	// akses map dengan cara namaVarMap["keyMap"]
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println(chicken["mei"])                // mei 0
}

// maps provide fast lookups and values that can retrieve, udpate, or delete with the help of keys
// maps is a reference to a hash table
// We can use any type in map index,
//Whereas arrays and slices can only use integers as indexes,
//a map can use any type for a keys (As long as values of that type can be compared using ==).