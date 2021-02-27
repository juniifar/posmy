package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Println("Getting PORT First")
	port, err := strconv.Atoi(os.Getenv("process.env.PORT"))
	if err == nil {
		fmt.Println("Error Getting PORT ??")
		log.Fatal(err)
	}

	fmt.Printf("Starting server at port %d\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
