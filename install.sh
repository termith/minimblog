apt-get install -y sqlite3 golang

mkdir -p $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:$(go env GOPATH)/bin
cd $GOPATH
go get -u github.com/termith/minimblog
cd src/github.com/termith/minimblog
go install

sqlite3 /etc/minimblog/blog.db engine/common/db/sql/up.sql
minimblog &