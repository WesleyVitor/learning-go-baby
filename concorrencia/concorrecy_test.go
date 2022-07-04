package concorrencia

import (
	"reflect"
	"testing"
	"time"
)


func mockVerificaWebsites(url string)bool{
	if url == "www.google.com.br" {
		return false
	}

	return true
}
func TestVerificaWebsites(t *testing.T){
	urls := []string{
		"www.google.com.br",
		"www.facebook.com",
		"www.youtube.com",
	}

	expected := map[string]bool{
		"www.google.com.br" : false,
		"www.facebook.com" : true,
		"www.youtube.com" : true,
	}

	resultado := VerificaWebsites(mockVerificaWebsites, urls)

	if !reflect.DeepEqual(resultado, expected){
		t.Errorf("got %v expected %v", resultado, expected)
	}
}

func slowStubVerificadorWebsite(url string)bool{
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificaWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i := 0; i < len(urls); i++ {
        urls[i] = "uma url"
    }

    for i := 0; i < b.N; i++ {
        VerificaWebsites(slowStubVerificadorWebsite, urls)
    }
}