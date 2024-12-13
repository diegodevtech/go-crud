FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/go-crud .
COPY /.env .

EXPOSE 8000

CMD ["./go-crud"]