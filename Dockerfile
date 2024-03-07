FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN make build-http


EXPOSE 3005

CMD ["./bin/http-server"]
