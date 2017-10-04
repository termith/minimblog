package post

const (
	CREATE_POST    = "INSERT INTO posts (title, text, created) VALUES ($1, $2, $3);"
	GET_LAST_POSTS = "SELECT title, text, created FROM posts LIMIT $1"
)

func GetLastPosts(count int) ([]*Post, error) {
	/*
		Return %count% of last posts from DB
	*/

}

func CreatePost(created int, title, text string) error {
	/*
		Create post with %title%, %text% created at %created% (Unix time)
	*/
	var err error
	return err
}
