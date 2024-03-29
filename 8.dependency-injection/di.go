package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}
