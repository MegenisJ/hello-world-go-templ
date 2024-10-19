package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
    component := hello("james")
    http.Handle("/", templ.Handler(component))
    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
