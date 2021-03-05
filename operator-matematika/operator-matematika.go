package main

import "fmt"

func operatorMatematika() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)
}

// cara singkat
func augmentedAssignments() {
	var i = 10
	i += 30 // 10 + 30
	fmt.Println(i)
}

func unaryOperator() {
	// ++
	var i = 2
	i++ // 2 + 1 = 3
	i++ // 3 + 1 = 4
	fmt.Println(i)

	// Boolean kebalikan
	var status bool = true
	if !status {
		fmt.Println("status kebalikan", status)
	} else {
		fmt.Println("status true", status)
	}
}

func main() {
	operatorMatematika()
}
