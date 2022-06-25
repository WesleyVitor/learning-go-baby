package main

const portuguese = "Portuguese"
const french = "French"
const englishPrefix = "Hello,"
const portuguesePrefix = "Olá,"
const frenchPrefix = "Bonjour,"

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language)
}

func greetingPrefix(language string)(prefix string){
	switch language{
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main(){

}