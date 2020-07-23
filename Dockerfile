FROM debian:latest

RUN apt-get update && \
    apt-get -y install gcc mono-mcs && \
    rm -rf /var/lib/apt/lists/*

FROM golang:latest


WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go get github.com/mattn/go-sqlite3
RUN go install -v ./...

ENV DB_DIALECT sqlite3
ENV DB_DATABASE test.db
EXPOSE 8081

CMD ["go","run","main.go"]