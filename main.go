package main

import "fmt"

func main() {
	var age int8 = 20
	var name string = "Muhammad Zaqi Al Quraisyi"
	university := "Asia Cyber University"
	println("Hello world")

	fmt.Println("Hello my name is ", name)
	fmt.Println("I am", age, "years old")
	fmt.Println("Currently i study of", university)
	{
		name := "Aaz"
		println("Overide variable value of name : ", name)
	}
	println("scope of variable in golang", name)
}
