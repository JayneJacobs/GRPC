package main

import (
	"fmt"
	"io"

	_ "github.com/golang/protobuf/protoc-gen-go"
)

func main() {
	fmt.Println("")
	io.Copy("/..")
}
