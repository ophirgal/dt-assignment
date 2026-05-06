# --- build stage ---
FROM golang:1.26.2 AS build

WORKDIR /app

# cache deps
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/main.go


# --- run stage ---
FROM alpine:latest

WORKDIR /app

# copy binary
COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]