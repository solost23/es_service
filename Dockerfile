FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn/

WORKDIR /build
COPY . .
RUN go mod tidy & go build -o app ./cmd/main.go
RUN apk add tzdata

FROM scratch

COPY --from=builder /build/app /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai

CMD ["/app"]