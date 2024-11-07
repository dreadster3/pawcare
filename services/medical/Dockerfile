ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /build

# Copy the Go Modules and download the dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o medical

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /build/medical .

ENV GIN_MODE=release

RUN adduser --no-create-home --disabled-password docker
RUN chown -R docker:docker /app

USER docker

EXPOSE 8080

CMD ["./medical"]

