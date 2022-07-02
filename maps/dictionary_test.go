package maps

import "testing"

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test":"This is just a test"}
	t.Run("key exists", func(t *testing.T) {
		
		got, _ := dictionary.Search("test")
		expected := "This is just a test"
		compareStrings(t,got,expected)
	})

	t.Run("key not exists", func(t *testing.T) {
		
		_, err := dictionary.Search("teto")
		if err == nil{
			t.Fatal("IT is need return a error")
		}
	})
	
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func compareStrings(t testing.TB, got, expected string){
	t.Helper()
	if got != expected{
		t.Errorf("got %s expected %s", got, expected)
	}

}

func TestUpdate(t *testing.T){
	t.Run("Key exists", func(t *testing.T) {
		key := "test"
		value := "Its just a test"
		new_value := "Its just another test"

		dictionary := Dictionary{key:value}

		err := dictionary.Update(key, new_value)

		assertError(t,err,nil)
		assertDefinition(t,dictionary, key, new_value)

	})

	t.Run("Key dont exists", func(t *testing.T) {
		key := "test"
		value := "Its just a test"
		

		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t,err,ErrWordDoesNotExist)
	})
}