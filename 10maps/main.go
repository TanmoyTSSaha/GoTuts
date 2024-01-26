package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Maps")
	languages := make(map[string]string)

	languages ["JS"] = "Javascript"
	languages ["TS"] = "Typescript"
	languages ["RB"] = "Ruby"
	languages ["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	
	delete(languages, "RB")
	
	fmt.Println("List of all languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For Key %v, Value is %v\n", key, value)
	}
}