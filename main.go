package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", makeHTTPhander(handleGetUser))
	http.ListenAndServe(":3000", nil)
}
