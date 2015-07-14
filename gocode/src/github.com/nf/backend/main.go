package main

import (
	"fmt"
//	"io/ioutil"
	"net/http"
//	"log"
)
/*
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Path[len("/view/"):]
	
	/*
	p, _ := loadPage(title)
	resp, err := http.Head("http://google.com/")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
    	log.Fatal(err)
	}

	fmt.Println(resp)
	fmt.Println(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)//how to respond to the frontend
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//fmt.Fprintf(w,command)

	switch command {
	    case "setIp1":
	    	fmt.Fprintf(w, "setIp1")
	    case "setIp2":
	    	fmt.Fprintf(w, "setIp2")
	    case "sumIp":
	    	fmt.Fprintf(w, "sumIp")
	    default:
	        fmt.Fprintf(w, "error has occurred")
    }
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}