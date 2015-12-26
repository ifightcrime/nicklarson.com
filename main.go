package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  //"fmt"
  "html/template"
)

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello world!")
    t, _ := template.ParseFiles("templates/home.tmpl")
    t.Execute(w, nil)
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":8001")
}
