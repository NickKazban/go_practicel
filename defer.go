////////12-14//////

/*package main

import "fmt"

func main() {
	fmt.Println("Good night")
	defer fmt.Println("world") /// викнується в останню чергу

	fmt.Println("hello")
	fmt.Println("hello hello")
	fmt.Println("hello hello hello")
}
*/

// //////13-14//////
package main

import "fmt"

func main() {
	fmt.Println("counting") // first is executed

	for i := 0; i < 10; i++ { //third is executed
		defer fmt.Println(i)
	}

	fmt.Println("done") //second is executed
}
