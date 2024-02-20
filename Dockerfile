# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

ARG APP_VERSION="undefined"
ARG BUILD_TIME="undefined"

WORKDIR /go/src/github.com/artarts36/regexlint

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o /go/bin/regexlint /go/src/github.com/artarts36/regexlint/cmd/main.go

######################################################

FROM alpine

COPY --from=builder /go/bin/regexlint /go/bin/regexlint

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.title="regexlint"
LABEL org.opencontainers.image.description="simple app for regex validation"
LABEL org.opencontainers.image.url="https://github.com/artarts36/regexlint"
LABEL org.opencontainers.image.source="https://github.com/artarts36/regexlint"
LABEL org.opencontainers.image.vendor="ArtARTs36"
LABEL org.opencontainers.image.version="$APP_VERSION"
LABEL org.opencontainers.image.created="$BUILD_TIME"
LABEL org.opencontainers.image.licenses="MIT"

EXPOSE 8080

ENTRYPOINT ["/go/bin/regexlint"]
