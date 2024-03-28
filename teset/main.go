package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	rel, _ := filepath.Rel("/a/data.json", "/a/temp.json")
	fmt.Println("Absolute path:", rel)
}
