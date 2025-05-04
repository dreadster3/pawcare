FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

WORKDIR /app/services/medical

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest

ARG HTTP_PORT=80
ARG GRPC_PORT=81

ENV HTTP_PORT=${HTTP_PORT}
ENV GRPC_PORT=${GRPC_PORT}

WORKDIR /app
COPY --from=builder /app/services/medical/server /app

EXPOSE ${HTTP_PORT}
EXPOSE ${GRPC_PORT}

USER nobody:nobody

ENTRYPOINT ["/app/server"]
