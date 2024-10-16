package main

import "fmt"

func main() {

// CONSTANT : NILAI TIDAK DAPAT DIUBAH LAGI
	// const firstName string = "Arya"
	// const lastName = "Perdana"

// MULTIPLE CONSTANT
	const (
		firstName string = "Arya"
		lastName         = "Perdana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	
	// error
	// firstName = "Tidak dapat diubah!"
	// lastName = "Tidak dapat diubah!"

}