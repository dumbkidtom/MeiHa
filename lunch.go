package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "encoding/json"
	// "regexp"
	// "net/url"
	// "html/template"
)

// controller REST API URL
var controller_url string = "http://localhost:5000/restaurants"

// business properties
type Business []struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Phone string `json:"phone"`
}

type Businesses struct {
	Business []struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
	}
	Total int `json:"total"`
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

	fmt.Fprintf(w, "response:%s", body)

	//t, _ := template.ParseFiles("card.gtpl")
	//t.Execute(os.Stdout, body)

}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/search", search_handler)
	http.ListenAndServe(":8080", nil)

}
