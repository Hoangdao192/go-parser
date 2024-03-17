package main

import (
	"joern-go/parser"
)

func main() {
	//projectDir := os.Args[1]
	//outDir := os.Args[2]
	//fmt.Println("Input project " + projectDir)
	//fmt.Println("Output directory " + outDir)
	projectDir := "/home/hoangdao/Workspace/Go/Test"
	outDir := "/home/hoangdao/Workspace/Go/GoParser/temp/test"
	parser.Parse(projectDir, outDir)

}
