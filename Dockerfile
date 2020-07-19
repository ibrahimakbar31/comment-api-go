FROM golang:1.14.4-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOCUSTOMENV=staging \
    GOCONFIGFILENAME=config

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]