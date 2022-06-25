package main

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t testing.TB,got string, want string){
		/*Helper Function to DRY of tests */
		t.Helper() // Allow to know what line was broke in realy test
		if got != want{
			t.Errorf("got %q want %q",got,want)
		}
	}


	t.Run("Should return hello to people", func(t *testing.T) {
		got := Hello("Bruna")
		want := "Hello,Bruna"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return Hello, World when pass a empty argument", func(t *testing.T){
		got := Hello("")
		want := "Hello,World"

		assertCorrectMessage(t, got, want)
	})

	
}