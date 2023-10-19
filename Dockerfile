FROM golang:1.21.3-alpine3.18

WORKDIR /app

RUN apk add --no-cache nodejs npm
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum package*.json ./
RUN go mod download
RUN npm i
COPY . .

CMD ["air", "-c", ".air.toml"]
