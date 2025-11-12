FROM golang:1.23-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
#ENV GOFLAGS="-mod=vendor"
WORKDIR /app

COPY . .

RUN go build -o app cmd/main.go cmd/must.go

FROM alpine:latest

RUN apk add --no-cache \
    curl \
    libc6-compat \
    gcompat \
    tzdata \
    ca-certificates


RUN addgroup -g 10001 app && \
    adduser -D -G app -h /app -u 10001 app
RUN update-ca-certificates
USER app
# path of application code at builder container and at this container must be the same for proper showing stacktrace at sentry
WORKDIR /app
EXPOSE 8000
ENTRYPOINT ["/app/app"]

# copying all code is required for proper showing stacktrace at sentry
COPY --from=builder /app/. /app/
