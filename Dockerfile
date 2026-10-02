FROM golang:1.24.9-alpine3.22 AS builder

WORKDIR /src

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/catalog-service \
    ./cmd/catalog-service

FROM scratch

COPY --from=builder /out/catalog-service /catalog-service

USER 65532:65532
EXPOSE 8080

ENTRYPOINT ["/catalog-service"]
