# Build stage
FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/main.go

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080
CMD [ "./api" ]