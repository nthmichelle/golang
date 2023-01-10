package main

import "fmt"

func main() {
	string := "Hello World!"

	fmt.Printf(string)

	// boolean
	boolVar := true

	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	//integer
	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	//float
	floatVar1 := float32(3.2)
	floatVar2 := float64(4.5)

	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T Value: %v\n", floatVar2, floatVar2)

	//bytes
	bytesVar := byte(255) //batas maksimal bytes 256
	fmt.Println("Type: %T Value: %v\n", bytesVar, bytesVar)
}
