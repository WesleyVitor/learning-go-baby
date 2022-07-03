package mocks

import (
	"fmt"
	"io"
)

type Sleeper interface{
	Sleep()
}


func CountDown(w io.Writer, s Sleeper){
	
	for i:=3; i > 0;i--{
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, "go!")
}