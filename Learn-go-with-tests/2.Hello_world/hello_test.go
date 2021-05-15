package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func (t testing.TB, got, want string) {
		t.Helper() // let test suite know this method is a helper method
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("saying hello to people", func (t *testing.T){
		got := Hello("Chris")
		want := "hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello Wrodl' when an empty string is passed", func (t *testing.T) {
		got := Hello("")
		want := "hello world"

		assertCorrectMessage(t, got, want)
	})
}