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
	//projectDir := "/home/hoangdao/Workspace/Go/GoParser/resources/test-project/gin"
	//outDir := "/home/hoangdao/Workspace/Go/GoParser/temp/gin"
	parser.Parse(projectDir, outDir)

}
