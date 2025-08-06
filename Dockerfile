# syntax=docker/dockerfile:1.5
FROM golang:1.24 as builder

WORKDIR /app
ARG TARGETARCH
ARG TARGETVARIANT
ADD . .
RUN CGO_ENABLED=0 GOARCH="${TARGETARCH}" GOARM="${TARGETVARIANT#v}" go build -ldflags="-w -s" -o dns-noerror main.go

FROM scratch
LABEL org.opencontainers.image.title="dns-noerror"
LABEL org.opencontainers.image.description="A dumb dns server that always returns NOERROR"
LABEL org.opencontainers.image.url="https://github.com/snowzach/dns-noerror"
LABEL org.opencontainers.image.source="https://github.com/snowzach/dns-noerror"
LABEL org.opencontainers.image.documentation="https://github.com/snowzach/dns-noerror#readme"

WORKDIR /app
COPY --from=builder /app/dns-noerror /app/dns-noerror
ENTRYPOINT ["/app/dns-noerror"]
