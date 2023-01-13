package main

import "fmt"

func main() {
	var chicken map [string]int
	chicken map [string]int {}
	
	chicken ["Januari"] = 50
	chicken ["Februari"] = 40 

	fmt.Println("Januari", chicken ["Januari"]) // januari 50
	fmt.Println("Mei", chicken ["Mei"]) // mei 0
}
