package main

import "fmt"

func main() {

	// var name string
	// name = "Arya Duta"
	// fmt.Println(name)

	// name = "Arya Perdana"
	// fmt.Println(name)

// TIDAK WAJIB MENYEBUT TIPE DATA VARIABLE JIKA LANGSUNG MENGINISIALISASI DATA PADA VARIABLENYA
	// var name = "Arya Duta"
	// fmt.Println(name)

	// name = "Arya Perdana"
	// fmt.Println(name)

// KATA KUNCI var TIDAK WAJIB!
// PERLU MENGGUNAKAN KATA KUNCI := SAAT MENGINISIALISASIKAN DATA PADA VARIABLE
	name := "Arya Duta" // banyak orang membuat variable seperti ini
	fmt.Println(name)

	name = "Arya Perdana"
	fmt.Println(name)
	
// DEKLARASI MULTIPLE VARIABLE
	var (
		firstName = "Arya"
		middleName = "Duta"
		lastName = "Perdana"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}