package hello

import (
	"crypto/rand"
	"fmt"
)

func Hello() string {
	return "Hello!\n"
}

func Foo() string {
	return "Foo Response\n"
}

func RandToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x\n", b)
}
