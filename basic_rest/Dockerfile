# syntax=docker/dockerfile:1

FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /basic_rest_server
EXPOSE 8080
CMD ["/basic_rest_server"]