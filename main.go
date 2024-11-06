package main

import (
	"fmt"
	"net/http"
)

func homePage(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe("", nil)

	if err != nil {
		fmt.Println(err)
	}
}
