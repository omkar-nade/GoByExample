package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the  weekend")
	default:
		fmt.Println("it's the weekdays")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a Bool")
		case int:
			fmt.Println("I'am a int")
		default:
			fmt.Printf("I don't know type %T/n ", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("Ram")
}
