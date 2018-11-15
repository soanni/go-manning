package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	urlString := "https://a b.com"
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("%#v\n", err)
		if urlError, ok := err.(*url.Error); ok {
			fmt.Printf("Op:", urlError.Op)
			fmt.Printf("URL:", urlError.URL)
			fmt.Printf("Err:", urlError.Err)
		}
		os.Exit(1)
	}
	fmt.Println(u)
}
