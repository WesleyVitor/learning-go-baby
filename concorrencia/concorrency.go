package concorrencia

import (
	"fmt"
	"runtime"
)

type VerificadorWebsite func(string) bool
type resultado struct{
	string
	bool
}
func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
    resultados := make(map[string]bool)
	channel := make(chan resultado)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GORA:", runtime.NumGoroutine())
    for _, url := range urls {
		go func(u string){
			channel <- resultado{u, vw(u)}
		}(url)
        
    }
	fmt.Println("GORD:", runtime.NumGoroutine())
	for i:=0;i<len(urls);i++{
		res := <- channel
		resultados[res.string] = res.bool
	}
    return resultados
}