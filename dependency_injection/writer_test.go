package dependencyinjection

import (
	"bytes"
	"testing"
)


func TestGreeting(t *testing.T){
	buffer := bytes.Buffer{}
	name := "Bruna"
	Greeting(&buffer, name)
	
	got := buffer.String()
	expected := "Olá, Bruna"

	if got != expected{
		t.Errorf("got %s expected %s", got, expected)
	}

}