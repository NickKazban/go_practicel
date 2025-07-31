// //1-26 Methods//

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())			//можна і так 	fmt.Println(Abs(v))
// }
// //3-26 Methods continued//

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// 4-26 Pointer receivers//
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(10)
// 	fmt.Println(v.Abs())
// }

// 5-26 Pointers and functions//

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v *Vertex, f float64) { // Ця функція масштабує вектор v на множник f. // Зверни увагу, що вона приймає вказівник *Vertex, тобто змінює значення v напряму.
// 	v.X = v.X * f //Зверни увагу, що вона приймає вказівник *Vertex, тобто змінює значення v напряму.
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10) // використання множинника для функції  Scale
// 	fmt.Println(Abs(v))
// }

// 6-26 Methods and pointer indirection//

// package main

// import "fmt"

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(2)
// 	ScaleFunc(&v, 10)

// 	p := &Vertex{4, 3}
// 	p.Scale(3)
// 	ScaleFunc(p, 8)

// 	fmt.Println(v, p)		 //4*3*8//3*3*8
// }

// 13-26 Nil interface values//
// package main

// import "fmt"

// type I interface {
// 	M()
// }

// func main() {
// 	var i I
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// 14-26 The empty interface//
// package main

// import "fmt"

// func main() {
// 	var i interface{} //викликає 1 раз
// 	describe(i)

// 	i = 42 //викликає 2 раз
// 	describe(i)

// 	i = "hello" //викликає 3 раз
// 	describe(i)

// }

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i) //викликає 3 рази
// }

// 16-26 Type switches//

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
