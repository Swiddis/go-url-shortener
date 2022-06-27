FROM golang:alpine AS builder
RUN apk add build-base
WORKDIR /src
COPY . .
RUN go mod download

FROM builder AS test
ENTRYPOINT [ "go", "test", "./..." ]

FROM builder AS build
RUN go build -o main

FROM alpine AS prod
WORKDIR /app
COPY --from=build /src/main .
EXPOSE 9808
ENTRYPOINT [ "./main" ]
