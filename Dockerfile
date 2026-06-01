# ---- Build Stage ----
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Bağımlılıkları kopyala ve indir
COPY go.mod ./
RUN go mod download

# Kaynak kodu kopyala ve derle
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ---- Run Stage ----
FROM alpine:3.19

WORKDIR /app

# Sadece derlenmiş binary'yi al (küçük image)
COPY --from=builder /app/server .

EXPOSE 8081

CMD ["./server"]
