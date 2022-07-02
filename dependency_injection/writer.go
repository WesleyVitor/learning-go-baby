package dependencyinjection

import (
	"fmt"
	"io"
)


func Greeting(w io.Writer ,name string){
	fmt.Fprint(w,"Olá, ", name)
}
