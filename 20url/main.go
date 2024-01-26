package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://itcorp.com:3000/it?dept=developer&company=appwrite";

func main() {
	fmt.Println("Welcome to handling URLs in GoLang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Schema: ",result.Scheme)
	fmt.Println("Host: ",result.Host)
	fmt.Println("HostName: ",result.Hostname())
	fmt.Println("Path: ",result.Path)
	fmt.Println("Port: ",result.Port())
	fmt.Println("RawQuery: ", string(result.RawQuery))
	fmt.Println("RawPath: ",result.RawPath)
	fmt.Println("RawFragment: ",result.RawFragment)
	
	qparams := result.Query()

	fmt.Printf("Query Params Type: %T\n",qparams)
	fmt.Println("Query Params: ",qparams["company"][0])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=tanmoy",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}