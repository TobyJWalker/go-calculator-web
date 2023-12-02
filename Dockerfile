FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

EXPOSE 8082

CMD ["/app/main"]