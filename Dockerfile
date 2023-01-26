FROM golang:1.16.5 as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app ./server

FROM alpine:3.14 as production
ENV DB_NAME=""
ENV GITHUB_TOKEN=""
ENV MongoURI=""
ENV COLLECTION_NAME=""
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder app .
# Exec built binary
CMD ./app
