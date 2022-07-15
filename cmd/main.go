package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	helloService "elastic.com/gorest"
)

func hello(w http.ResponseWriter, r *http.Request) {
	msg, error := helloService.Hello("there")
	if error != nil {
		fmt.Fprintf(w, "error")
	}
	w.Write([]byte(msg))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
