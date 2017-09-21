# sudo docker build -t minimblog .
# sudo docker run -d -p 8080:8080 --rm --name minimblog_backend minimblog
FROM golang:1.8

WORKDIR /minimblog

ENV SRC_DIR=/go/src/github.com/termith/minimblog/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o minimblog; cp minimblog /minimblog/
ENTRYPOINT ["./minimblog", "--config", "/go/src/github.com/termith/minimblog/resources/config.json"]