FROM golang:1.16-alpine
WORKDIR /app
COPY *.go ./

COPY go.mod ./

RUN go build -o /http-server
EXPOSE 8080
CMD ["/http-server"]