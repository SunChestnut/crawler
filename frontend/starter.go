package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.SearchResultHandler{})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
