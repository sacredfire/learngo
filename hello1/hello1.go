package main

import "fmt"

const fuckersName string = "Idiot Fuck"

// func main() {
// 	// fmt.Println("go" + "lang")

// 	// fmt.Println("1+1 =", 1+1)
// 	// fmt.Println("7.0/3.0 =", 7.0/3.0)

// 	// fmt.Println(true && false)
// 	// fmt.Println(true || false)
// 	// fmt.Println(!true)

// 	// var a string = "initial"
// 	// fmt.Println(a)

// 	// var b, c int = 2, 4
// 	// fmt.Println(b, c)

// 	// var d = true
// 	// fmt.Println(d)

// 	// var e int
// 	// fmt.Println(e)

// 	// f := "short"
// 	// fmt.Println(f)

// 	// fmt.Println(32132 * 42452)

// 	// fmt.Println(len(a))

// 	// fmt.Println((true && false) || (false && true) || !(false && false))

// 	// var x string = "Hello World"
// 	// fmt.Println(x)

// 	// var y string
// 	// y = "Hi, there!"
// 	// fmt.Println(y)
// 	// y = y + "second"
// 	// fmt.Println(y)

// 	// s := "fuck you all"
// 	// fmt.Println(s)

// 	// fmt.Println("Fuckers name is", fuckersName)

// 	// fuckersName := "Fucky Fuck"
// 	// fmt.Println("Fuckers name is", fuckersName)
// }

// func main() {
// 	fmt.Print("Enter a number: ")
// 	var input float64
// 	fmt.Scanf("%f", &input)

// 	output := input * 2

// 	fmt.Println(output)
// }

func main() {
	fmt.Print("Enter a temp in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9

	fmt.Println("You temp in Celcius is", output)
}
