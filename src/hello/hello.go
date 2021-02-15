package main
import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const spanish = "Spanish"
func Hello(name string, language string) string{
	if name == "" {
		name = "world"
	}
	prefix := englishPrefix
	if language == spanish {
		prefix = spanishPrefix
	}
	return prefix + name
}

func main(){
	fmt.Println(Hello("world", ""))
}
