FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV HOST_HTTP=0.0.0.0
ENV HOST_PORT=3005

ENV HTTP_DETAILED_ERROR=false

ENV DATABASE_URL=postgresql://fibo:fibo@postgres:5432/fibo
ENV ACCESS_TOKEN_EXPIRES_TTL=180    
ENV ACCESS_TOKEN_SECRET=secret

RUN make build-http


EXPOSE 3005

CMD ["./bin/http-server"]
