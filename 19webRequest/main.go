package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://appwrite.io/"

func main() {
	fmt.Println("Welcome to Web Requests in GoLang")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of Type: %T\n", response)

	defer response.Body.Close() // CALLER'S RESPONSIBILTY TO CLOSE THE CONNECTION.

	databytes, err := io.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)

	writeFiles("./itcorp.html", content)

}

func writeFiles(filename string, content string) {
	file, err := os.Create(filename);

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Print("length is: ", length)

	file.Close()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}