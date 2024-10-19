package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
    "github.com/MegenisJ/hello-world-go-templ/Pages"
)

func main() {
    component := pages.index()
    http.Handle("/", templ.Handler(component))
    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
