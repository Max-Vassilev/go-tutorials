package main

import "fmt"

func main() {
	// Array:
	var intArr [3]int32 = [3]int32{1, 2, 3}

	// Print element from the array:
	fmt.Println(intArr[0])

	// Print the memory location of an element from the array- using "&"":
	fmt.Println(&intArr[0])

	// Slice:
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println("Before append - len:", len(intSlice), "cap:", cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println("After append  - len:", len(intSlice), "cap:", cap(intSlice))

	// Better to set capacity early to avoid memory reallocation when appending.
	intSlice2 := append(make([]int32, 0, 8), 4, 5, 6)
	fmt.Println("Before append - len:", len(intSlice2), "cap:", cap(intSlice2))

	// Map:

}
