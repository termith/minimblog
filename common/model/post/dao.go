package post

const (
	CREATE_POST = "INSERT INTO posts (title, text, created) VALUES ($1, $2, $3); SELECT last_insert_rowid() from posts;"
)
