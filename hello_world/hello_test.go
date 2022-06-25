package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Bruna")
	want := "Hello,Bruna"

	if got != want{
		t.Errorf("got %q want %q",got, want)
	}
}