FROM golang:alpine AS build
RUN apk add --update --no-cache ca-certificates git
COPY . $GOPATH/src/github.com/Ahton89/r2s/
WORKDIR $GOPATH/src/github.com/Ahton89/r2s/
RUN go get -d -v
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/r2s
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/r2s /bin/
ENTRYPOINT ["r2s"]