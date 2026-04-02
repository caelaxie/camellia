package abbrcase

import (
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
)

// Analyzer reports identifier declarations that use all-caps abbreviations.
var Analyzer = &analysis.Analyzer{
	Name: "abbrcase",
	Doc:  "reports identifier declarations that use all-caps abbreviations instead of camel-case abbreviations",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	projectFiles := projectFileSet(pass.Files, pass.Fset)

	for ident, obj := range pass.TypesInfo.Defs {
		if ident == nil || obj == nil || ident.Name == "_" {
			continue
		}

		if _, ok := obj.(*types.PkgName); ok {
			continue
		}

		if !projectFiles.contains(pass.Fset.File(ident.Pos())) {
			continue
		}

		suggestion, changed := SuggestedName(ident.Name)
		if !changed {
			continue
		}

		pass.Reportf(ident.Pos(), "identifier %q should use camel-case abbreviations: %q", ident.Name, suggestion)
	}

	return nil, nil
}

type fileSet map[string]struct{}

func projectFileSet(files []*ast.File, fset *token.FileSet) fileSet {
	names := make(fileSet, len(files))
	for _, file := range files {
		if file == nil {
			continue
		}

		if tokFile := fset.File(file.Pos()); tokFile != nil {
			names[filepath.Clean(tokFile.Name())] = struct{}{}
		}
	}

	return names
}

func (f fileSet) contains(tokFile *token.File) bool {
	if tokFile == nil {
		return false
	}

	_, ok := f[filepath.Clean(tokFile.Name())]
	return ok
}
