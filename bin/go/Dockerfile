FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
RUN apk add bash
WORKDIR /go/src/app
ADD ./go .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./cmd/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080