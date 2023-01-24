FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o out .

FROM alpine:latest

COPY --from=builder /app/out /app/api

EXPOSE 8080

CMD ["/app/api"]