package main

import "fmt"
import "strings"

/*
Penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga dry (kependekan dari don't repeat yourself),
tak perlu menuliskan banyak kode yang kegunaannya berkali-kali,
cukup sekali saja lalu panggil sesuai kebutuhan.
cara baca-tulis : declare-name-assign-type-value-address

:= declares and assigns, = just assigns
It's useful when you don't want to fill up your code with type or struct declarations.

*/

func main() {
	var names = []string{"John", "Wick", "Budianto"} //composite literal type string
	// names := []string{"John", "Wick", "Budianto"}
	// variabel bernama names bertipe slice-string bervalue {"John", "Wick", "Budianto"}

	printMessage("Hi", names) //2 dipanggil fungsi printMessage beserta masing2 nilai parameter
	// message valuenya = "halo" dan arr valuenya = names = slice string "[]string" dari variabel names yg isinya "john","wick"
}

func printMessage(message string, arr []string) { //1 dibuat fungsi dengan 2 parameter : 1.message dan 2.arr
	var fullName = strings.Join(arr, " ") //2 dibuat variabel fullname dengan
	// fullName := strings.Join(arr, " ")
	fmt.Println(message, fullName)
}
