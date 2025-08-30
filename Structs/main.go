package main

import "fmt"

type Vertex struct {
	// X int
	// Y int
	X, Y int
}

// make fields of structs in capital letters do we can access in another ex : Post, Content
// we can use the structs for parsing the front-end json file into structs field

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 Y:0 is implicit
	p  = &Vertex{1, 2} // has a type *Vertex
)

func main() {
	// fmt.Println(Vertex{1, 2})
	// v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v.X)

	// p := &v
	// p.X = 6
	// fmt.Println(v.X)

	fmt.Println(v1, v2, v3, p)

}
