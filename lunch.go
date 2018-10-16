package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//  "net/url"
	//  "html/template"
)

// controller REST API URL
var controller_url string = "http://localhost:5000/restaurants"

// restaurant properties
type Restaurant struct {
	Name  string
	URL   string
	Phone string
}

func index_handler(w http.ResponseWriter, r *http.Request) {

	// t,_ := template.ParseFiles("input.gtpl")
	// t.Execute(w, nil)
	http.ServeFile(w, r, "input.gtpl")

}

func search_handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	myzip := r.Form.Get("zip")
	myradius := r.Form.Get("radius")

	// call Controller REST API for results
	url := fmt.Sprintf("%s?zip=%s&radius=%s", controller_url, myzip, myradius)

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "body:%s", body)

	// convert response to JSON
	//var bytes = []byte(body)
	// Unmarshal string into structs.
	var restaurants []Restaurant
	json.Unmarshal(body, &restaurants)

	fmt.Fprintf(w, "response:%s", restaurants)
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/search", search_handler)
	http.ListenAndServe(":8080", nil)

}
