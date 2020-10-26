FROM golang:latest

WORKDIR /go/src/github.com/3nt3/urlshortener
COPY . . 

RUN go get -v -d ./...
RUN go install -v ./...

CMD ["urlshortener"]
