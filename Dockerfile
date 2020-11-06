FROM golang:1.14

WORKDIR /go/src/space-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GIN_MODE=release
CMD ["go", "run", "./src/main.go"]