# Middleware-server

### Init

> go mod init example/"project-name"

### Dependencies
- echo (go get github.com/labstack/echo/v4, go get github.com/labstack/echo/v4/middleware@v4.10.0)
- backoff (go get github.com/cenkalti/backoff/v4)
- crdb (go get github.com/cockroachdb/cockroach-go/v2/crdb)
- pq (go get github.com/lib/pq)

### go mod tidy

> 쓰지 않는 packages는 다 날려준다.