package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
    pages "megenisj/hello-world/pages"
)

func main() {
    component := pages.Index()
    http.Handle("/", templ.Handler(component))
    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
