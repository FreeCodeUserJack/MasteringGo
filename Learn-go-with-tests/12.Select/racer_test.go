package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func TestRacer(t *testing.T) {
	t.Run("check result", func(t *testing.T) {
		slowServer := makeDelayedServer(1 * time.Second)
		fastServer := makeDelayedServer(0 * time.Second)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect error but got one %v", err)
		}
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(1 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 500 * time.Millisecond)

		if err == nil {
			t.Errorf("expected error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}