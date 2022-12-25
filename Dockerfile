FROM golang:1.18.5-buster
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

COPY ./. .

RUN go mod download
CMD ["air", "-c", ".air.toml"]