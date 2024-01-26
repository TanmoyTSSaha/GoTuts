package main

import "fmt"

func main() {
	fmt.Println("Struct in GoLang")

	tanmoy := User{"Tanmoy", "tanmoy@go.dev", true, 20}

	fmt.Printf("User details: %+v",tanmoy)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
} 