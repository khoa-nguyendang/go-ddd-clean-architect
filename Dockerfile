FROM golang:alpine as builder
COPY . /app
WORKDIR  /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build  -tags musl -ldflags='-s -w '  -o api-server ./cmd/...

FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /app/api-server /api-server