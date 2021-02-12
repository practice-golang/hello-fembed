package main // import "hello.wasm"

import (
	_ "embed"
	"fmt"
)

//go:embed "files/hello.txt"
var b []byte

func main() {
	fmt.Println("Hello, " + string(b) + "!")
}
