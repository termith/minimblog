package post

import (
	"time"
)

type Post struct {
	Title string
	Text string
	TimeCreated time.Time
}