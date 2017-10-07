package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/termith/minimblog/common/model/post"
	"github.com/termith/minimblog/engine/common/logging"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Handles post creation
	*/
	switch r.Method {
	case "POST":
		r.ParseForm()
		title := r.Form.Get("title")
		body := r.Form.Get("body")
		created := time.Now().Unix()
		err := post.CreatePost(created, title, body)
		if err != nil {
			logging.Error("Error while creating post")
			logging.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		logging.Warn("Only POST method supported")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Handles getting last posts from DB
	*/
	switch r.Method {
	case "GET":
		posts, err := post.GetPosts()
		if err != nil {
			logging.Error("Error while getting posts from DB")
			logging.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		logging.Warn("Only GET method supported")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, no %s here yet!", r.URL.Path[1:])
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
