package main

import (
	"fmt"
	"joern-go/parser"
	"os"
)

func main() {
	projectDir := os.Args[1]
	outDir := os.Args[2]
	fmt.Println("Input project " + projectDir)
	fmt.Println("Output directory " + outDir)
	//projectDir := "/home/hoangdao/Workspace/Go/GoParser/resources/test-project/simple"
	//outDir := "/home/hoangdao/Workspace/Go/GoParser/temp/simple"
	parser.Parse(projectDir, outDir)

}
