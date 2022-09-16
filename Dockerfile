FROM golang:1.18.3-alpine

WORKDIR /app

RUN mkdir bin
RUN mkdir src

WORKDIR /app/src

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

COPY . .
RUN go build -o /app/bin/main main.go

WORKDIR /app

RUN rm -r /app/src/

WORKDIR /app/bin

COPY .env .

CMD ["./main", "server"]