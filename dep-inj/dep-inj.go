package inj

import (
	"io"
	"fmt"
	// "bytes"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}