package main

import "fmt"

type transFxType func(int) int

func main() {
	numbers := []int{2, 4, 6, 8, 9}
	moreNumbers := []int{3, 6, 9, 12, 15}

	fmt.Println("************************** Using Transformer Numbers*************************************")
	doubles := transFormNumbers(&numbers, doubleNumber)
	tripples := transFormNumbers(&numbers, trippleNumber)
	fmt.Println(doubles)
	fmt.Println(tripples)

	fmt.Println("************************** Using Transformer Numbers With Anonymous Funx *************************************")
	doubles2 := transFormNumbers(&numbers, func(number int) int { return number * 2 })
	tripples2 := transFormNumbers(&numbers, func(number int) int { return number * 3 })
	fmt.Println(doubles2)
	fmt.Println(tripples2)

	fmt.Println("************************** Using Factory Pattern to create Transformer Numbers With Anonymous Funx and closure *************************************")
	funxDouble := createTransformer(2)
	funxTripper := createTransformer(3)
	doubles3 := transFormNumbers(&numbers, funxDouble)
	tripples3 := transFormNumbers(&numbers, funxTripper)
	fmt.Println(doubles3)
	fmt.Println(tripples3)

	fmt.Println("************************** Using Transformer Numbers and Get Transform Funcsx*************************************")
	doubles4 := transFormNumbers(&numbers, getTransformFunc(&numbers))
	tripples4 := transFormNumbers(&moreNumbers, getTransformFunc(&moreNumbers))
	fmt.Println(doubles4)
	fmt.Println(tripples4)

	fmt.Println("************************** Using Recursion Funcx*************************************")
	fact := factorial(6)
	fmt.Println(fact)

	fmt.Println("************************** Using VariadicFunc Funcx*************************************")
	someArray := []int{5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(VariadicFunc(4, someArray...)) // can also accept array, just unpack it with ...
	fmt.Println(VariadicFunc(4, 5, 6, 7, 8, 9))
	fmt.Println(VariadicFunc(4, 5, 6, 7, 8))
	fmt.Println(VariadicFunc(4, 5, 6, 7))
	fmt.Println(VariadicFunc(4, 5, 6))
	fmt.Println(VariadicFunc(4, 5))
	fmt.Println(VariadicFunc(4))
}

// Note if transFx type is very long, you can give it an alias
// type transFxType func(int) int, and use transFxType in the func below :)
func transFormNumbers(list *[]int, transFx func(int) int) []int {
	response := []int{}
	for _, value := range *list {
		response = append(response, transFx(value))
	}
	return response
}

// double the parameter and return it
func doubleNumber(number int) int {
	return number * 2
}

// tripple the parameter and return it
func trippleNumber(number int) int {
	return number * 3
}

// Tranformer decision Maker, to use func as a return value with aliased parameter type
func getTransformFunc(numbers *[]int) transFxType {
	if (*numbers)[0] == 2 {
		return doubleNumber
	} else {
		return trippleNumber
	}
}

// Using closure with anonymous function in factory pattern and return it
func createTransformer(factor int) transFxType {
	return func(number int) int {
		return number * factor
	}
}

// Using Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// Using VariadicFunc
func VariadicFunc(starter int, numbers ...int) int {
	sum := starter
	for _, value := range numbers {
		sum += value
	}
	return sum
}
