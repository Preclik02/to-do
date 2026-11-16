package main

import(

	"fmt"
	"strings"

)

func main() {

	text := "hello;world"

	parts := strings.Split(text, ";")

	fmt.Println(parts)
}
