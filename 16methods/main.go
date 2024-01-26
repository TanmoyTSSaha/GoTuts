package main

import "fmt"

func main() {
	fmt.Println("Struct in GoLang")

	tanmoy := User{"Tanmoy", "tanmoy@go.dev", true, 20}

	fmt.Printf("User details: %+v\n",tanmoy)

	tanmoy.GetStatus()
	tanmoy.NewEmail("myemail@go.dev")
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail(email string) {
	u.Email = email
	fmt.Println("New email is: ", u.Email)
}