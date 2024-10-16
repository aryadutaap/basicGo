package main

import "fmt"

func main() {
// Type Declaration : membuat ulang tipe data baru dari tipe data yang sudah ada
	type NoKTP string

	var ktpDuta NoKTP = "1111111111"

	var contoh string = "2222222222"
	var ContohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpDuta)
	fmt.Println(ContohKtp)
	fmt.Println(NoKTP("3333333333"))
}