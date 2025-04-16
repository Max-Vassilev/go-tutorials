package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var maxNum16 int16 = 32767
	fmt.Println(maxNum16)
	var maxNum32 int32 = 2147483647
	fmt.Println(maxNum32)
	var maxNum64 = 9223372036854775807
	fmt.Println(maxNum64)

	var exampleFloat32 float32 = 3.1415927
	fmt.Println(exampleFloat32)
	var exampleFloat64 float64 = 3.141592653589793
	fmt.Println(exampleFloat64)

	var myString string = "Hello" + " " + "world!"
	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString(myString))

	var myBool bool = true
	fmt.Println(myBool)

	const myConstStr string = "Const string"
	fmt.Println(myConstStr)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
