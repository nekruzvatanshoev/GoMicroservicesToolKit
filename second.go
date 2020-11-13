package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()


	router.HandleFunc("/", helloWorld)
	fmt.Println("Connecting...")
	if err := http.ListenAndServe(":80", router); err != nil {
		fmt.Println(err)
	}
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}