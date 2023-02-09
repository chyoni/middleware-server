# Middleware-server

### Init

> go mod init example/"project-name"

### Dependencies
- echo (go get github.com/labstack/echo/v4, go get github.com/labstack/echo/v4/middleware@v4.10.0)
- backoff (go get github.com/cenkalti/backoff/v4)
- crdb (go get github.com/cockroachdb/cockroach-go/v2/crdb)
- pq (go get github.com/lib/pq)
- dockertest (go get -u github.com/ory/dockertest/v3)
    - Dockertest를 쓰는 이유는 거의 모든 서비스들이 데이터베이스와 의사소통할텐데 그런 것들을 테스트하기 위해 mocking 하는게 너무 귀찮고 힘드니까 이를 해결해주기 위해 실제 데이터베이스를 만들고 테스트 후에 바로 없애버리는 
    그런 테스팅을 하게 된다. 그리고 이를 위해 Docker가 최고의 시스템이기 때문에 Docker를 이용해서 테스트하고 테스트 다 끝나면 Container를 내려버리는 뭐 이런 로직을 사용하는 테스트 라이브러리인가보다.

- require (go get github.com/stretchr/testify/require)
### go mod tidy

> 쓰지 않는 packages는 다 날려준다.


### Go routines

- 동시에 여러개를 실행할 수 있는 병렬적 실행 방법
- go routine과 통신하기 위해서는 channel이 필요하다.