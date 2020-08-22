## Build

FROM golang:1.15.0-alpine AS builder
WORKDIR /home/ai/go/src/github.com/classified5/gocker
COPY . .
RUN go build -o binary main.go

## Distribute

FROM alpine:latest
RUN mkdir /app && mkdir gocker
WORKDIR /gocker
EXPOSE 9090
COPY --from=builder /home/ai/go/src/github.com/classified5/gocker/binary /app
CMD /app/binary