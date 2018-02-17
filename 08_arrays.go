package gobyexample

import "fmt"

// PrintSomeArrays - prints some arrays
func PrintSomeArrays() {
	// declare zero-valued array
	var a [5]int
	fmt.Println("emp:", a)

	// set and get values in an array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// get length of an array
	fmt.Println("len:", len(a))

	// declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("decl:", b)

	// two-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
