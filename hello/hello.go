package hello

import (
	"fmt"

	stringcase "github.com/reiver/go-stringcase"
)

func SayHello() {
	fmt.Printf("hello, world\n")
	s := stringcase.ToUpperCase("Hello world")
	fmt.Printf("%s\n", s)
}
