package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run(`existing key`, func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run(`non-existing key`, func(t *testing.T) {
		_, err := dictionary.Search("foo")
		want := ErrorWordNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run(`new word`, func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run(`existing word`, func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run(`update existing word`, func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run(`update new non-existing word`, func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newWord := "test new"
		newDefinition := "new definition"
		err := dictionary.Update(newWord, newDefinition)

		assertError(t, err, UpdateErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run(`delete an existing word`, func(t *testing.T) {
		word := "test"
		definition := "test definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrorWordNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertNoError(t, err)
	assertString(t, got, definition)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
