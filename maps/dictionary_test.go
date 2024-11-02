package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		err := d.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		newDefinition := "new definition"
		err := d.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{}
		err := d.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	d := Dictionary{word: "test definition"}
	d.Delete(word)
	_, err := d.Search(word)

	assertError(t, err, ErrNotFound)
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error %v, want %v", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
