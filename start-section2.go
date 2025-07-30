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
// package main

// import "fmt"

// type Vertex struct {
// 	X, Y int
// }

// var (
// 	v1 = Vertex{1, 2}  // has type Vertex
// 	v2 = Vertex{X: 1}  // Y:0 is implicit
// 	v3 = Vertex{}      // X:0 and Y:0
// 	p  = &Vertex{1, 2} // has type *Vertex
// )

//	func main() {
//		fmt.Println(v1)
//		fmt.Println(v2)
//		fmt.Println(v3)
//		fmt.Println(p)
//		fmt.Println(v1, p, v2, v3)
//	}
//
// ////6-27 Arrays/////
// package main

// import "fmt"

// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes)
// }
// ////7-27 Slices/////

// package main

// import "fmt"

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

//		var s []int = primes[1:4]
//		fmt.Println(s)
//	}
//
// ////8-27 Slices are like references to arrays////
// package main

// import "fmt"

// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	a := names[0:2]
// 	b := names[1:3]
// 	fmt.Println(a, b)

//		b[0] = "XXX"
//		fmt.Println(a, b)
//		fmt.Println(names)
//	}
//
// ////9-27 Slice literals////
// package main

// import "fmt"

// func main() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q)

// 	r := []bool{true, false, true, true, false, true}
// 	fmt.Println(r)

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
// }
// ////10-27 SSlice defaults////

// package main

// import "fmt"

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}

// 	s = s[0:2]
// 	fmt.Println(s)

// 	s = s[1:4]
// 	fmt.Println(s)

// 	s = s[0:2]
// 	fmt.Println(s)

// 	s = s[:2]
// 	fmt.Println(s)

// 	s = s[1:]
// 	fmt.Println(s)

// 	s = s[0:2]
// 	fmt.Println(s)
// }
// ////11-27 Slice length and capacity////

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
