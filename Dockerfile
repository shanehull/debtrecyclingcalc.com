FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN apk update && apk upgrade && apk add --no-cache ca-certificates git
RUN update-ca-certificates

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN GO_ENABLED=0 GOOS=linux go build \
    -ldflags "debtrecyclingcalc.com/internal/buildinfo.GitTag=$(git describe --tags)" \
    -o ./bin/main ./cmd/

FROM scratch

ENV ALLOWED_ORIGIN="*"
ENV SERVER_HOST="0.0.0.0"

COPY --from=builder /app/bin/main ./bin/main
COPY --from=builder /app/static ./static

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./bin/main"]
