# 1) Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./app

# 2) Final minimal image
FROM scratch
# copy only the statically-linked binary (including embedded assets)
COPY --from=builder /src/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
