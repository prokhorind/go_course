//go:build ignore

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, _ := filepath.Glob("router/*.go")

	outputFile := "router/generated_routes.go"
	if err := os.MkdirAll("router", 0755); err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(outputFile)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	out.WriteString("package router\n//AUTO generated code do not modify\nfunc init() {\n")

	fset := token.NewFileSet()
	for _, file := range files {
		if strings.Contains(file, "generated_routes.go") {
			continue
		}

		node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		for _, decl := range node.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				continue
			}
			for _, comment := range fn.Doc.List {
				if strings.HasPrefix(comment.Text, "// @route ") {
					route := strings.TrimSpace(strings.TrimPrefix(comment.Text, "// @route "))
					out.WriteString("\tregisterRoute(\"" + route + "\", " + fn.Name.Name + ")\n")
				}
			}
		}
	}

	out.WriteString("}\n")
}
