# syntax=docker/dockerfile:1
FROM golang:1.19.2-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-testproject
EXPOSE 8080
CMD [ "/docker-testproject" ]