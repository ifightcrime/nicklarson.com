package main

import (
	"github.com/codegangsta/negroni"

	"net/http"
	//"fmt"
	"html/template"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello world!")

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		t, _ := template.ParseFiles("templates/home.tmpl")
		t.Execute(w, nil)
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8001")
}
