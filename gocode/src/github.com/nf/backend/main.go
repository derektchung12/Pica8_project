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

var ip1 string
var ip2 string



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
	    	setIp1(w,r)
	    	//fmt.Fprintf(w, "setIp1")

	    case "setIp2":
	    	setIp2(w,r)
	    	//fmt.Fprintf(w, "setIp2")
	    case "sumIp":
	    	sumIp(w,r)
	    	//fmt.Fprintf(w, "sumIp")
	    default:
	        fmt.Fprintf(w, "error has occurred")
    }
}

func setIp1(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    //fmt.Println(r.Form)
    //fmt.Println(r.FormValue("ip1value"))

    ip1 = r.FormValue("ip1value")

    ip1 +=".1"
    //ip1 = "setIP1"
    fmt.Fprintf(w, ip1)
}

func setIp2(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	ip2 = r.FormValue("ip1value")
	ip2 += ".1"
	fmt.Fprintf(w, ip2)
}

func sumIp(w http.ResponseWriter, r *http.Request){
	var stringsum string= ip1+ip2
	fmt.Fprintf(w, stringsum) 
}



func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}