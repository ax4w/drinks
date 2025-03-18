package main

import (
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(Index()))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err.Error())
	}
}
