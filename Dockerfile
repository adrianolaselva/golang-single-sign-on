FROM golang:1.13.11 as builder

WORKDIR /app/

COPY . .

RUN go get -d -v ./... \
    && GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o ./oauth2

FROM golang:1.13.11-stretch

WORKDIR /app

COPY --from=builder ./app/public ./public
COPY --from=builder ./app/doc ./doc
COPY --from=builder ./app/oauth2 .

ENTRYPOINT ["./oauth2"]