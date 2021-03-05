package main

import "fmt"

// type
type Int32 int32
type Int64 int64
type NoTaksi string
type Status bool

// contoh di struct
type Car struct {
	Name    string
	NoTaksi NoTaksi
	Active  Status
}

func main() {
	var noTaksiGolden NoTaksi = "B 2584 HY"
	var berangkatStatus Status = true
	fmt.Println(noTaksiGolden)
	fmt.Println(berangkatStatus)
}
