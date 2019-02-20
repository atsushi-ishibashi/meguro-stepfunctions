FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/meguro-stepfunctions/fargate
RUN \
  apk update && \
  rm -rf /var/cache/apk/* && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go build -ldflags "-s -w" -o meguro .

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/meguro-stepfunctions/fargate/meguro /usr/local/bin/
ENTRYPOINT ["meguro"]
