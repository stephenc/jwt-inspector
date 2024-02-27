FROM golang:alpine as app

RUN addgroup -S golangdocker \
 && adduser -S -u 10000 -g golangdocker golangdocker
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app /go/bin/jwt-inspector /jwt-inspector
COPY --from=app /etc/passwd /etc/passwd
USER golangdocker
ENTRYPOINT ["/jwt-inspector"]