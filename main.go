package main

import "fmt" // fmt use for formatted I/O

func main() {
	// String
	var nameOne string = "Alice"
	var nameTwo = "Bob"
	var nameThree string

	//update
	nameOne = "Charlie"
	nameThree = "David"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Eve" // := syntax means declare and initialize with string type, Go infers the type from the value on the right
	fmt.Println(nameFour)

	//init

	var ageOne int = 25
	var ageTwo = 30
	ageThree := 35

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numFive int8 = 100
	var numSix int8 = -128
	var numSeven uint16 = 256

	fmt.Println(numFive, numSix, numSeven)

	// float
	var priceOne float32 = 19.99
	var priceTwo float64 = 29.99
	priceThree := 39.99

	fmt.Println(priceOne, priceTwo, priceThree)

}
