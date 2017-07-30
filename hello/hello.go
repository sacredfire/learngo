package main

import "fmt"

// func main() {
// 	fmt.Printf("hello, world\n")

// 	var c int

// 	a := 10
// 	b := 9

// 	c = a + b

// 	fmt.Println(c)

// 	fmt.Println("1 + 1 =", 1+1)

// 	fmt.Println(len("Hello World"))
// 	fmt.Println("Hello World"[1])
// 	fmt.Println("Hello " + "World!")

// 	fmt.Println(true && true)
// 	fmt.Println(true && false)
// 	fmt.Println(true || true)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)

// }

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "even")
// 		} else {
// 			fmt.Println(i, "odd")
// 		}
// 	}

// 	for i := 1; i <= 10; i++ {
// 		switch i {
// 		case 0:
// 			fmt.Println("zero")
// 		case 1:
// 			fmt.Println("one")
// 		case 2:
// 			fmt.Println("two")
// 		case 3:
// 			fmt.Println("three")
// 		case 4:
// 			fmt.Println("four")
// 		case 5:
// 			fmt.Println("five")
// 		case 6:
// 			fmt.Println("six")
// 		case 7:
// 			fmt.Println("seven")
// 		case 8:
// 			fmt.Println("eight")
// 		case 9:
// 			fmt.Println("nine")
// 		case 10:
// 			fmt.Println("ten")
// 		}
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// 	fmt.Println(x[4])
// }

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83
// 	var total float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	fmt.Println(total / float64(len(x)))
// }

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

// some commeny just to test commit via vsstudio
