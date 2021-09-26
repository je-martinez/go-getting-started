package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//Types
	p := Person{name: "Jose", age: 25}
	pList := []Person{}
	pList = append(pList, p)
	pList[0].name = "Hola"
	fmt.Println(p)
	fmt.Println(pList)
	i := 5
	notModifiedRef(&i)
	println(i)
}

func notModifiedRef(x *int) {
	*x++
}
