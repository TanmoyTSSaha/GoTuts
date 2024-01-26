package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the Functions in GoLang")
	greeting()
	result := adder(3,5);
	fmt.Println("Result: ", result)
	
	var proResult int = proAdder(4, 5,1,7,0,5)
	fmt.Println("Pro-result: ", proResult)
}

func greeting() {
	fmt.Println("Namastay from GoLang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo;
}

func proAdder (values ...int) int {
	var total int = 0;

	for _, value := range values {
		total += value;
	}

	return total;
}