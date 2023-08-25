package main

import (
	"fmt"
	"net/http"

	"github.com/aniketDinda/mongoapi/router"
)

func main() {
	fmt.Println("Mongo DB API")
	r := router.Router()
	fmt.Println("Server is getting Started ...")

	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at 4000")
}
