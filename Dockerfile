FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 4040

CMD ["air", "-c", ".air.toml"]
