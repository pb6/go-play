FROM golang:1.13-alpine
WORKDIR /go/src/
RUN apk --no-cache add git && \
    go get github.com/lib/pq
COPY envtester.go .
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o app envtester.go

FROM scratch
COPY --from=0 /go/src/app /app
EXPOSE 8080
USER 1000
ENTRYPOINT ["/app"]
