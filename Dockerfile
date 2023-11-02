FROM golang:1.21.3-alpine3.18

WORKDIR /app

RUN apk add --no-cache nodejs npm

COPY go.mod go.sum package*.json ./
RUN go mod download
RUN npm i
COPY . .
RUN npm run build:css
RUN go build -o ./bin/main ./cmd/website/main.go

CMD "./bin/main"
