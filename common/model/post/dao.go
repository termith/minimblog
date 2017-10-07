package post

import (
	"github.com/termith/minimblog/common/db"
)

const (
	CREATE = "INSERT INTO posts (title, text, created) VALUES ($1, $2, $3);"
	LIST   = "SELECT title, text, created FROM posts LIMIT"
)

func GetPosts(count int) ([]*Post, error) {
	/*
		Return posts from DB
	*/
	rows, err := db.Connection().Query(LIST)
	var posts []*Post
	return posts, err
}

func CreatePost(created int64, title, text string) error {
	/*
		Create post with %title%, %text% created at %created% (Unix time)
	*/
	_, err := db.Connection().Query(CREATE, title, text, created)
	return err
}
