FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

RUN go install github.com/a-h/templ/cmd/templ@latest

ARG GIT_TAG="unknown"

RUN GO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X debtrecyclingcalc.com/internal/buildinfo.GitTag=${GIT_TAG}"\
    -o ./bin/main ./cmd/

FROM scratch

ENV ALLOWED_ORIGIN="*"
ENV SERVER_HOST="127.0.0.1"

COPY --from=builder /app/bin/main ./bin/main
COPY --from=builder /app/static ./static

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./bin/main"]
