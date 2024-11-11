FROM golang:1.16
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go build -o main .
CMD ["/app/main"]
