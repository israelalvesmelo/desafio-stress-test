FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stress-test ./cmd/cli

FROM scratch
COPY --from=builder /app/stress-test .

ENTRYPOINT ["./stress-test"]
