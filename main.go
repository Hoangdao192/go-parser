package main

import "joern-go/parser"

func main() {

	// projectDir := os.Args[1]
	projectDir := "/home/hoangdao/Workspace/Go/GoParser"
	outDir := "/home/hoangdao/Workspace/Go/GoParser/temp"
	parser.Parse(projectDir, outDir)
}
