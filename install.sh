apt-get install -y sqlite3 golang

mkdir -p $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:$(go env GOPATH)/bin
cd $GOPATH
go get -u github.com/termith/minimblog
cd src/github.com/termith/minimblog
go install

mkdir -p /etc/minimblog
cp resources/config.json /etc/minimblog/congig.json

sqlite3 /etc/minimblog/blog.db engine/common/db/sql/up.sql
minimblog /etc/minimblog/config.json &