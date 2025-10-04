FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -o api

FROM alpine:latest

WORKDIR /app/

COPY --from=build /app/api .

CMD ["./api"]