package pkg

import (
	"fmt"
	"runtime"
)

// To print system's go version from terminal
func Version() {
	ver := runtime.Version()
	fmt.Println(ver)
}
