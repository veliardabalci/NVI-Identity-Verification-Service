# 1. Go imajını kullan
FROM golang:1.22.4

# 2. Çalışma dizinini ayarla
WORKDIR /app

# 3. Gerekli dosyaları kopyala
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 4. Uygulamayı derle
RUN go build -o main .

# 6. Uygulamayı çalıştır
CMD ["./main"]
