# syntax=docker/dockerfile:1
FROM golang:1.19.2-alpine


WORKDIR /app
EXPOSE 8080

COPY . .

RUN go mod download
RUN go build -o ./app cmd/api/main.go

CMD [ "./app" ]