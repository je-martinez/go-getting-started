package main

import (
	"errors"
	"fmt"
	"math"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	//First Method
	var x int = 3
	var y int = 2
	var sum int = x + y
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("Sum", sum)
	//Second Method Infer it
	x1 := 3
	y2 := 5
	sum2 := x1 + y2
	fmt.Println("x1", x1)
	fmt.Println("y2", y2)
	fmt.Println("sum2", sum2)
	//Conditions
	if y > y2 {
		fmt.Println(y, "higher than", y2)
	} else if y < y2 {
		fmt.Println(y, "less than", y2)
	} else if y == y2 {
		fmt.Println(y, "equal than", y2)
	}
	//Arrays With Fix N of Elements
	var a [5]int
	a[2] = 5
	b := [5]int{5, 2, 5, 4}
	fmt.Println(a)
	fmt.Println(b)

	//Arrays with Indeterminate n of elements
	c := []int{}
	c = append(c, 25)
	fmt.Println(c)

	//Map
	vertices := make(map[string]int)
	vertices["key_1"] = 2
	vertices["key_2"] = 3
	vertices["key_3"] = 25
	fmt.Println(vertices)
	delete(vertices, "key_3")
	fmt.Println(vertices)

	//Loops with starting point to end point
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Loop and array or something like that
	arr := []int{5, 4, 3, 2, 1}
	for index, value := range arr {
		fmt.Println("Index", index, "Value", value)
	}

	myMap := make(map[string]int)
	myMap["A"] = 1
	myMap["B"] = 2
	myMap["C"] = 3

	for key, value := range myMap {
		fmt.Println("Key", key, "Value", value)
	}
	result := sum3(5, 5)
	result2, err := sqrt(-3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	fmt.Println("Sum Two Numbers With A Funciton", result)
}

//Normal Function
func sum3(x int, y int) int {
	return x + y
}

//Function with multiple returns
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
