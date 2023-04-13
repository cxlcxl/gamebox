FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

#COPY templates /app/templates
COPY assets/images/game-icons /app/assets/images/game-icons
#COPY config/games.ini /app/config/games.ini
#COPY config/web.example.yaml /app/config/web.yaml
RUN go build -ldflags="-s -w" -o /app/web/gamebox web/main.go


FROM scratch

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
#ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/web/gamebox /app/web/gamebox
COPY --from=builder /app/assets /app/assets

CMD ["./web/gamebox"]
