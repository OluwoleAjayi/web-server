package main

import "fmt"

func main() {
	arr := [...]int{3, 42, 2}

	// for i := 1; i < 100; i++ {
	// 	fmt.Println(fizzBuzz(i))
	// }
	fmt.Print(arr)
}

// func fizzBuzz(num int) string {
// 	switch {
// 	case num%3 == 0:
// 		fmt.Println("Fizz")
// 	case num%5 == 0:
// 		fmt.Println("Buzz")
// 	case num%15 == 0:
// 		fmt.Println("FizzBuzz")
// 	}

// 	return strconv.Itoa(num)

// }
