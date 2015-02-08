package main

import (
	"fmt"

	"github.com/hnakamur/w32version"
)

func main() {
	v, err := w32version.GetVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Windows %s", v)
}
