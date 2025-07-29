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

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	X := "first"
	Y := "second"

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.X, v.Y)
	fmt.Println(X, Y)
}
