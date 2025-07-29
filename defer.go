package main

import "fmt"

func main() {
	fmt.Println("Good night")
	defer fmt.Println("world") /// викнується в останню чергу

	fmt.Println("hello")
	fmt.Println("hello hello")
	fmt.Println("hello hello hello")
}
