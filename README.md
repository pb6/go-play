Features:
  /ip - get public external ip
  /env - dump container environment
  /pg - check postgresql connectivity

To build:

  CGO_ENABLED=0 go build -a -ldflags '-s' envtester.go
  docker build .

To run:
  docker run -tiP <image>

For postgresql to be successful, pass env variable connStr:
  postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full

See https://godoc.org/github.com/lib/pq for more.
