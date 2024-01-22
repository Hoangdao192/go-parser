package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"joern-go/util"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func isGoFile(info fs.FileInfo) bool {
	return !info.IsDir() &&
		strings.LastIndex(info.Name(), ".go")+len(".go") == len(info.Name())
}

func main() {

	// projectDir := os.Args[1]
	projectDir := "/home/hoangdao/Workspace/Go/GoParser"
	outDir := "/home/hoangdao/Workspace/Go/GoParser/temp"

	_, openFileErr := os.OpenFile(outDir, os.O_RDONLY, os.ModePerm)
	if openFileErr != nil {
		if mkdirErr := os.MkdirAll(outDir, os.ModePerm); mkdirErr != nil {
			log.Fatal(mkdirErr)
		}
	}

	err := filepath.Walk(projectDir, func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if isGoFile(info) {
			os.MkdirAll(path.Dir(filepath), os.ModePerm)

			fileContent, readErr := util.ReadFile(filepath)
			if readErr == nil {
				parsedFile, parseErr := parser.ParseFile(token.NewFileSet(), info.Name(),
					fileContent, parser.AllErrors)
				if parseErr == nil {
					jsonData, jsonErr := json.Marshal(parsedFile)
					if jsonErr == nil {
						saveFilePath := outDir + "/" + filepath[strings.Index(
							filepath, projectDir):] + ".json"
						saveFile, openFileErr := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
							os.ModePerm)
						if openFileErr != nil {
							log.Fatal(openFileErr)
						}
						saveFile.Write(jsonData)
					}
				} else {
					log.Fatal(parseErr)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.DeclarationErrors)
	if err != nil {
		log.Fatal(err)
	}
	var v visitor
	ast.Walk(v, f)
}

type visitor int

func (v visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), node)
	return v + 1
}
