apt-get install -y sqlite3 golang

mkdir -p $HOME/go
export GOPATH=$HOME/go
cd $GOPATH

go get -u github.com/termith/minimblog
cd src/github.com/termith/minimblog
go build minimblog.go -o minimblog

sqlite3 resources/blog.db engine/common/db/sql/up.sql
./minimblog --config config.json &
