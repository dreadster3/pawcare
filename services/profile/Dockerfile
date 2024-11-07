ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /build

# Copy the source code
COPY . .

WORKDIR /build/services/profile

# Download the dependencies
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o profile 

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /build/services/profile/profile .

ENV GIN_MODE=release

RUN adduser --no-create-home --disabled-password docker
RUN chown -R docker:docker /app

USER docker

EXPOSE 8080

CMD ["./profile"]

