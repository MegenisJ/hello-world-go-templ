package main

import (
	"fmt"
	components "megenisj/hello-world/Components"
	pages "megenisj/hello-world/pages"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
    component := pages.Index()
    http.Handle("/", templ.Handler(component))
    button := components.Button()
    http.Handle("/clicked", templ.Handler(button))
    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
