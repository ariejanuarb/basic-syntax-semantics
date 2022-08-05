package main

import "fmt"

func main() {
	_ = "belajar Golang"   // variabel underscore adalah predefined, jadi tidak perlu penggunakan :=
	_ = "Golang itu mudah" // biasa digunakan untuk menampung nilai balik fungsi yg tidak digunakan

	// untuk pengisian nilai multi variabel yang dilakukan dengan metode type inference,
	//boleh di dalamnya terdapat variabel underscore.
	name, _ := "john", "wick"

	// Data yang sudah masuk variabel tersebut akan hilang.
	// Ibaratkan variabel underscore seperti blackhole, o
	// objek apapun yang masuk ke dalamnya, akan terjebak selamanya di-dalam singularity dan tidak akan bisa keluar

	fmt.Println(name)
}
