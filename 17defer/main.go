package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer in GoLang")
	
	defer fmt.Println("Hello World!")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}