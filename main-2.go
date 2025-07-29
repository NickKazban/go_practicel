////9-14 ///////

/*package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
*/

////////10-14///

/*package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(time.Now())    //Real time now

}*/

//// Різниця між днями (субота-сьогодні)/////

/*package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	// Вычисляем разницу между сегодняшним днем и субботой
	daysUntilSaturday := (time.Saturday - today + 7) % 7

	// Используем switch для вывода результата
	switch daysUntilSaturday {
	case 0:
		fmt.Println("Today.")
	case 1:
		fmt.Println("Tomorrow.")
	case 2:
		fmt.Println("In two days.")
	default:
		fmt.Printf("Go to Saturday  to  %d days. Good luck man", daysUntilSaturday)
	}
}*/

