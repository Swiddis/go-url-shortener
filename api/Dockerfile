FROM golang:alpine AS builder
RUN apk add build-base
WORKDIR /src
RUN go mod init github.com/Swiddis/go-url-shortener
RUN go get github.com/go-redis/redis/v8
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/itchyny/base58-go/cmd/base58
RUN go get github.com/stretchr/testify
RUN go mod download
COPY . .

FROM builder AS test
ENTRYPOINT [ "go", "test", "./..." ]

FROM builder AS build
RUN go build -o main

FROM alpine AS prod
WORKDIR /app
COPY --from=build /src/main .
EXPOSE 9808
ENTRYPOINT [ "./main" ]

