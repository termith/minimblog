package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type post struct {
	Title     string
	Body      string
	Author    string
	Timestamp time.Time
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	myPost := post{"My first Title", "What\nis\nup ", "Dmitry Demidov", time.Now()}
	bytesResponse, _ := json.Marshal(myPost)
	w.Header().Set("Content-type", "application/json")
	w.Write(bytesResponse)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, no %s here yet!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/post/", postHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
