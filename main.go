package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening to port 5000")

	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
