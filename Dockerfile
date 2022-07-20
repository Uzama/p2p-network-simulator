##
## Build
##
FROM golang:1.18.4-alpine3.16 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o p2p-network-simulator

##
## Deploy
##
FROM alpine:latest

WORKDIR /

COPY --from=build /app/p2p-network-simulator /p2p-network-simulator

EXPOSE 8080

ENTRYPOINT ["./p2p-network-simulator"]