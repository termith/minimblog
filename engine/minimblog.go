package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
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

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	basePath := "../frontend"
	if r.URL.Path == "/" {
		basePath = basePath + "/index.html"
	}
	requested := filepath.Join(basePath, r.URL.Path)
	info, err := os.Stat(requested)
	if err != nil {
		if os.IsNotExist(err) || info.IsDir() {
			http.NotFound(w, r)
			return
		}
		http.Error(w, http.StatusText(500), 500)
	}

	http.ServeFile(w, r, requested)
	//fmt.Fprintf(w, "Hi there, no %s here yet!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/post/", postHandler)
	http.HandleFunc("/new/", newPostHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
