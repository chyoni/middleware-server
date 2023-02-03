package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

func TestResponds(t *testing.T) {
	// ! Docker로 테스트를 수행할건데 그 pool을 만드는 statement
	pool, err := dockertest.NewPool("")
	// ! NoError()는 err가 nil임을 assert하는 function
	require.NoError(t, err, "colud not connect to Docker")

	// ! middleware-server-mock:latest 라는 컨테이너를 실행하는 statement
	resource, err := pool.Run("middleware-server-mock", "latest", []string{})
	require.NoError(t, err, "could not start container")

	// ! Cleanup은 callback함수를 받는데 그 함수는 이 Test와 그 하위 Test들이 모두 종료된 후 실행된다.
	// ! 약간 defer라고 생각해도 될 거 같다.
	// ! 그리고 가장 마지막에 추가된 Cleanup이 가장 먼저 호출된다.
	t.Cleanup(func() {
		// ! Purge는 컨테이너와 컨테이너에 연결된 volume을 지운다.
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response
	// ! Retry()는 여러번 실행하는거야 될 때까지 아무리해도 계속 안되면 에러를 던진다.
	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8081/tcp"), "/"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "Http status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	require.Contains(t, string(body), "Hello", "does not respond with Hello?")
}
