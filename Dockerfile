FROM golang:alpine AS builder

WORKDIR /src
RUN go mod init github.com/Swiddis/go-url-shortener
RUN go get github.com/go-redis/redis/v8
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/itchyny/base58-go/cmd/base58
RUN go mod download
COPY . .
RUN go build -o main

FROM alpine

WORKDIR /app
COPY --from=builder /src/main .

EXPOSE 9808
CMD [ "./main" ]

