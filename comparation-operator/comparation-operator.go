package main

import "fmt"

func main() {

	var name1 = "Golden"
	var name2 = "Bird"

	var result bool = name1 == name2
	fmt.Println(result, "result")

	var val1 = 500
	var val2 = 700

	fmt.Println(val1 > val2)  // false
	fmt.Println(val1 < val2)  // true
	fmt.Println(val1 == val2) // false
	fmt.Println(val1 != val2) // true
}
