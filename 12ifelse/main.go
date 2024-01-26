package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else in GoLang")

	loggedInUser := 10
	var result string

	if loggedInUser > 10 {
		result = "Noo not this!"
	} else if loggedInUser < 10 {
		result = "Not even this one!"
	} else {
		result = "This is exectly what she said!"
	}

	fmt.Println(result)

	if num:=10; num%2==0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}
}