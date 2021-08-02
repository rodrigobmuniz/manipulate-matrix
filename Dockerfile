FROM golang:1.16.5 as builder
WORKDIR /app
COPY . .
RUN go build -o manipulate-matrix main.go
CMD ./manipulate-matrix
