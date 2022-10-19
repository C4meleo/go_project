package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	matches, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		os.Exit(2)
	}
	for _, file := range matches {
		target, _ := os.Readlink(file)
		if len(target) > 0 {
			fmt.Printf("%+v\n", target)
		}
	}
}
