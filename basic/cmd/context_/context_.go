package context_

import (
	"fmt"
	"net/http"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

func server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()

		// ↓これの Done() を解読するところから

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
		time.Now().Add(time.Second)
	}
}
