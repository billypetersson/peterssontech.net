// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
