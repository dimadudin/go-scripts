package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		if i%5 != 0 && i%3 != 0 {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
