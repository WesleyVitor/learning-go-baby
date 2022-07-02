package main

import (
	"fmt"
	di "learning-go-baby/dependency_injection"
	"log"
	"net/http"
)

func MyGreeterHandler(res http.ResponseWriter, req *http.Request){
	di.Greeting(res, "World")
}

func main(){
	
	err := http.ListenAndServe(":5000",http.HandlerFunc(MyGreeterHandler))
	if err != nil{
		log.Fatal("Servidor Caiu!")
	}

	fmt.Println("Servidor Rodando...")
}