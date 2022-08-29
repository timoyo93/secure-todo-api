FROM golang:1.18-alpine3.16

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init -g cmd/web/main.go
RUN go build -o /usr/local/bin/app cmd/web/main.go
EXPOSE 8080
CMD ["app"]