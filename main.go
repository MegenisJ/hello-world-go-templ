package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	components "megenisj/hello-world/Components"
	pages "megenisj/hello-world/pages"
)

func main() {
    component := pages.Index()
    http.Handle("/", templ.Handler(component))

    button := components.Button()
    http.Handle("/clicked", templ.Handler(button))

    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
