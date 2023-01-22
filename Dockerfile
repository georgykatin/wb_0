FROM golang:latest AS BUILDER
RUN mkdir /go/src/wb
WORKDIR /go/src/wb
COPY . .
RUN go mod download


WORKDIR /go/src/wb/main/subscriber
RUN go build -o subscriber main2.go
CMD ["./subscriber"]

WORKDIR /go/src/wb/main/publisher
RUN go build -o publisher main.go
CMD ["./publisher"]