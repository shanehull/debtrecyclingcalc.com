FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN apk update && apk upgrade && apk add --no-cache ca-certificates curl libstdc++ gcompat
RUN update-ca-certificates

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN ARCH=$(uname -m | sed 's/x86_64/x64/;s/aarch64/arm64/') && \
    curl -sLO "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-${ARCH}-musl" && \
    mv "tailwindcss-linux-${ARCH}-musl" tailwindcss && \
    chmod +x tailwindcss && \
    ls -lh tailwindcss

RUN templ generate

RUN ./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

ARG GIT_TAG="unknown"

RUN CGO_ENABLED=0 GOOS=linux go build \
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
