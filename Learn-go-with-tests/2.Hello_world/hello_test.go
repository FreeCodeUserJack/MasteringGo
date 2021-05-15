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
		got := Hello("Chris", "")
		want := "hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello Wrodl' when an empty string is passed", func (t *testing.T) {
		got := Hello("", "")
		want := "hello world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func (t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "hola Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func (t *testing.T) {
		got := Hello("Janice", "French")
		want := "bonjour Janice"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Southern", func (t *testing.T) {
		got := Hello("Pardner", "Southern")
		want := "howdy Pardner"

		assertCorrectMessage(t, got, want)
	})
}