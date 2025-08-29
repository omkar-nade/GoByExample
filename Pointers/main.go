package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// i, j := 42, 2710

	// p := &i
	// fmt.Println(*p)
	// *p = 345
	// fmt.Println(*p)

	// p = &j
	// *p = *p / 37
	// fmt.Println(*p)

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)

	fmt.Println("zeroval:", i)
	zeroptr(&i)

	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
