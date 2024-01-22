package main

import (
	"fmt"
	"path"
)

type Test struct {
	number int
}

func main() {
	
	outDir := "/home/hoangdao/Workspace/Go/GoParser/hello.go"
	fmt.Println(path.Dir(outDir))
}

func printTest(test Test) {
	fmt.Printf("%p\n", &test)
}
