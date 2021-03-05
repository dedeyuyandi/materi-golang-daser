package main

import "fmt"

func main() {
	const (
		firstName string = "Golden"
		lastName         = "Bird"
	)

	// // error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidal Bisa diubah"

	fmt.Println(firstName)
	fmt.Println(lastName)
}
