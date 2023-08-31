FROM golang:1.21-alpine AS builder


ADD cmd /go/src/application/cmd
ADD internal /go/src/application/internal
ADD pkg /go/src/application/pkg
ADD go.mod /go/src/application
ADD go.sum /go/src/application
WORKDIR /go/src/application/cmd

RUN go build -o /app


# Final stage
FROM alpine

WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
