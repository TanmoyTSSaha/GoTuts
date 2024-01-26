package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("Files in GoLang")

	content := "This needs to go in a file - www.google.com";

	filename := "./myLcoGoFile.txt";

	file, err := os.Create(filename)

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Print("length is: ", length)

	defer file.Close()

	readFile(filename)
}


func readFile(filename string) {
	databyte, err := os.ReadFile(filename);

	checkNilError(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))
}


func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}