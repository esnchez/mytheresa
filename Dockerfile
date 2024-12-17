# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsufix cgo -o api cmd/api/*.go

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/api .
# COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
# RUN chmod +x /app/wait-for.sh
# RUN chmod +x /app/start.sh            
# COPY db/migrations ./migrations

EXPOSE 8080
CMD [ "./api" ]
# ENTRYPOINT [ "/app/start.sh" ]