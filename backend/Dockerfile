# syntax=docker/dockerfile:1

FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

FROM golang:1.22-alpine3.19 AS runner

WORKDIR /app

COPY --from=builder /app/app ./

EXPOSE 8000

CMD ["/app/app"]

FROM golang:1.22-bookworm AS develop

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0

CMD ["air"]


# FROM golang:1.22-bookworm AS debugger
# WORKDIR /app
# COPY . .
# RUN go get github.com/go-delve/delve/cmd/dlv
# EXPOSE 3001 40000
# ENTRYPOINT ["dlv", "debug", "--listen=:40000", "--headless", "--accept-multiclient", "--continue", "--api-version=2"]

# FROM alpine:latest AS production
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /app .
# EXPOSE 3001
# ENTRYPOINT ["./main"]
