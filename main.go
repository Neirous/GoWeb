package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " This is home page")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " This is about page")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Bytes written:", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
