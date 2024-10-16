package main

import "fmt"

func main() {

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // number overflow : batas int16 = 32767 -> lebih 1, memutar ke -32768

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // >32767   -->   -32768

	var name = "Arya Duta" 
	var e uint8 = name[0] // char to byte
	var eString = string(e) // byte to string

	fmt.Println(name)
	fmt.Println(e) // output byte 
	fmt.Println(eString) // output string 
}