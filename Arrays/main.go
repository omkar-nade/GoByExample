package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emt:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println(a[4])
	fmt.Println("length of an array is :", len(a))

	// inti and declare an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", c)

	d := [...]int{100, 3: 300, 400, 500}
	fmt.Println("idx :", d)
	fmt.Println("idx:", len(d))
}
