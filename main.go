package main

import (
	"fmt"
	"log"
	"net/http"
)

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
