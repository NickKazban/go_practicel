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

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) { // Ця функція масштабує вектор v на множник f. // Зверни увагу, що вона приймає вказівник *Vertex, тобто змінює значення v напряму.
	v.X = v.X * f //Зверни увагу, що вона приймає вказівник *Vertex, тобто змінює значення v напряму.
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // використання множинника для функції  Scale
	fmt.Println(Abs(v))
}
