package main

import (
	"joern-go/parser"
)

func main() {

	// projectDir := os.Args[1]
	projectDir := "/home/hoangdao/Workspace/Go/GoParser/resources/test-project/simple"
	outDir := "/home/hoangdao/Workspace/Go/GoParser/temp/simple"
	parser.Parse(projectDir, outDir)

}
