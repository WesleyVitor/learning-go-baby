package interate
const countLimit = 5
func Repeat(character string) (string, int){
	word := ""

	for i:=0;i<countLimit;i++{
		word += character
	}

	return word, countLimit
}