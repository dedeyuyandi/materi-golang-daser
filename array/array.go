package main

import "fmt"

func array() {
	var arrString [3]string

	arrString[0] = "Blue Bird"
	arrString[1] = "Golden Bird"
	arrString[2] = "Lainnya"

	fmt.Println(arrString[0])
	fmt.Println(arrString[1])
	fmt.Println(arrString[2])

	// loop
	for i, val := range arrString {
		fmt.Println("index: ", i, "values:", val)
	}

	// panjang / jumlah array [3]
	fmt.Println(len(arrString))
}

// [Spike] Set Up Local Environment
// golang, bloomRPC, postman, proto, 

func arrayLangung() {
	var arrInt = []int{
		90,
		95,
		80,
	}

	// loop
	for i, val := range arrInt {
		fmt.Println("index: ", i, "values:", val)
	}

}

func main() {
	arrayLangung()
}
