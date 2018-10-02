package main

import (
  "fmt"
  "net/http"
//  "html/template"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

 // t,_ := template.ParseFiles("input.gtpl")
 // t.Execute(w, nil)
  http.ServeFile(w, r, "input.gtpl")

}

func search_handler(w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  myzip := r.Form["zip"]
  myradius := r.Form["radius"]

  fmt.Fprintf(w, "Received: zip %s, radius %s\n", myzip, myradius)
  fmt.Fprintf(w, "  calling Eric for data...")
}

func main() {

   http.HandleFunc("/", index_handler)
   http.HandleFunc("/search", search_handler)
   http.ListenAndServe(":8080", nil)

}

