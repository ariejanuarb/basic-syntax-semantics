package main

import "fmt"

func main() {
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// jika key nya string, maka akan dicetak berdasarkan urutan alfabet key-string nya
	fmt.Println("chicken1 :", chicken1)

	var chicken2 = map[string]int{
		"januari":  60,
		"februari": 40,
	}
	fmt.Println("chicken2 :", chicken2)

	var chicken3 = map[string]int{
		"januar": 90,
		"arie":   40,
	}
	fmt.Println("chicken3 :", chicken3)

	var chicken4 = make(map[int]string)
	chicken4[10] = "arie"
	chicken4[1] = "januar"
	// jika key nya int, akan di print berdasarkan urutan angkanya
	fmt.Println("chicken4 :", chicken4)

	var chicken5 = *new(map[int]string)
	fmt.Println("chicken5 :", chicken5)
}
