FROM golang:1.9-alpine as builder
WORKDIR /go/src/csportal-server
COPY . .
RUN apk update && apk add git && apk add ca-certificates && apk add --no-cache curl
RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/csportal-server

FROM scratch
COPY --from=builder /go/src/csportal-server/bin/csportal-server ./csportal-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

CMD ["./csportal-server"]
