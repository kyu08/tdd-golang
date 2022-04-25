package context_

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello"
	t.Run("return a response", func(t *testing.T) {
		svr := server(&SpyStore{data, false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel  work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := server(store)

		// リクエストを生成
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// cancellingCtx, cancel を取得
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// 5ms後に cancel を実行する
		time.AfterFunc(5*time.Microsecond, cancel)
		// 新しいリクエストを生成
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		// リクエストを実行
		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})

	// 疑問: cancel ってなに？現実世界でいうどういうユースケースなのかがいまいち理解できていない
}
