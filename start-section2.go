///////1-27/////

/*package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer- 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i. - 21 нове значення через показчик "і"

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j - 73
}
*/

//////2-27 A struct is a collection of fields./////

/*package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
*/
//////3-27 Struct Fields/////

// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	X := "first"
// 	Y := "second"

//		v := Vertex{1, 2}
//		v.X = 4
//		fmt.Println(v.X)
//		fmt.Println(v.X, v.Y)
//		fmt.Println(X, Y)
//	}
//
// ////5-27 Pointers to structs/////
// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

//	func main() {
//		v := Vertex{1, 2}
//		p := &v
//		p.X = 1e9
//		fmt.Println(v)
//	}
//
// ////5-27 Struct Literals/////
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p)
	fmt.Println(v1, p, v2, v3)
}
