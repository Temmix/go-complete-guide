package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"Running", "Jumping", "clapping"}
	fmt.Println()
	fmt.Println("****************Simple Array***************************")
	arrayWith3Hobbies(hobbies)

	fmt.Println()
	fmt.Println("****************Display Different Array***************************")
	displayDifferentArray(hobbies)

	fmt.Println()
	fmt.Println("****************Two Slice Array***************************")
	newHobbies := twoSlice(hobbies)

	fmt.Println()
	fmt.Println("****************Reslice Array***************************")
	reSlice(hobbies, newHobbies)

	fmt.Println()
	fmt.Println("****************Dynamic Array***************************")
	dynamic := dynamicArray()

	fmt.Println()
	fmt.Println()
	fmt.Println("****************Add to Dynamic Array***************************")
	addToDynamicArray(dynamic)

	fmt.Println()
	fmt.Println("****************Product Array***************************")
	productArray()

	fmt.Println()
	fmt.Println("****************Append Array With Array***************************")
	appendArrayWithArray()

	fmt.Println()
	fmt.Println("****************Append Array With Array***************************")
	list := usingMake()
	fmt.Println(list)

	fmt.Println()
	fmt.Println("****************Looping Array***************************")
	LoopingArray(list)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.

func arrayWith3Hobbies(hobbies [3]string) {
	fmt.Println(hobbies)
}

// 2) Also output more data about that array:
//   - The first element (standalone)
//   - The second and third element combined as a new list
func displayDifferentArray(hobbies [3]string) {
	fmt.Println(hobbies[0])
	newHobbies := hobbies[1:]
	fmt.Println(newHobbies)
}

//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
func twoSlice(hobbies [3]string) []string {
	firstSlice := hobbies[0:2]
	secondSlice := hobbies[:2]
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	return firstSlice
}

//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
func reSlice(hobbies [3]string, newHobbies []string) {
	//newHobbies[0] = hobbies[1]
	//newHobbies[1] = hobbies[2]
	// it can still access the underlying array which is the main hobbies,also its capacity is 3 like the that of hobbies
	newHobbies = newHobbies[1:3]
	fmt.Println(newHobbies)
}

// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
func dynamicArray() []string {
	list := []string{"Learning", "Reading"}
	fmt.Print(list)
	return list
}

// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
func addToDynamicArray(list []string) {
	list[1] = "Researching"
	list = append(list, "Playing")
	fmt.Println(list)
}

//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.
func productArray() {
	list := []Product{
		{title: "First Product", id: 1, price: 29.99},
		{title: "Second Product", id: 2, price: 59.99},
	}

	fmt.Println(list)
	list = append(list, Product{title: "Third Product", id: 3, price: 79.99})
	fmt.Println(list)
}

// 8 . Append array to any Array with 3 dots ...
func appendArrayWithArray() {
	prices := []float64{12.99, 13.99, 14.99}
	fmt.Println(prices)

	// add more array to an array
	discountedPrices := []float64{2.99, 3.99, 4.99}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}

// 9. Using make to create an array
func usingMake() []string {
	list := make([]string, 5)
	list[0] = "Testing 1"
	list[1] = "Testing 2"
	list[2] = "Testing 3"
	list[3] = "Testing 4"
	list[4] = "Testing 5"

	list = append(list, "Testing 6")
	return list
}

// 10. Using looping
func LoopingArray(list []string) {
	for index, value := range list {
		fmt.Printf(`Index: %v, Value: %v`, index, value)
		fmt.Println()
	}
}
