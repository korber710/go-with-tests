package mocking

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	fmt.Fprint(writer, "3")
}
