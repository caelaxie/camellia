package camellia

import (
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
)

// Analyzer reports identifier declarations that use all-caps abbreviations.
var Analyzer = mustNewAnalyzer(Config{})

// Config defines analyzer behavior.
type Config struct {
	Exclude []string
}

type analyzerRunner struct {
	matcher *excludeMatcher
}

// NewAnalyzer builds a camellia analyzer from configuration.
func NewAnalyzer(config Config) (*analysis.Analyzer, error) {
	matcher, err := newExcludeMatcher(config.Exclude)
	if err != nil {
		return nil, err
	}

	runner := analyzerRunner{matcher: matcher}

	return &analysis.Analyzer{
		Name: "camellia",
		Doc:  "reports identifier declarations that use all-caps abbreviations instead of camel-case abbreviations",
		Run:  runner.run,
	}, nil
}

func mustNewAnalyzer(config Config) *analysis.Analyzer {
	analyzer, err := NewAnalyzer(config)
	if err != nil {
		panic(err)
	}

	return analyzer
}

func (r analyzerRunner) run(pass *analysis.Pass) (any, error) {
	projectFiles := projectFileSet(pass.Files, pass.Fset, r.matcher)

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

func projectFileSet(files []*ast.File, fset *token.FileSet, matcher *excludeMatcher) fileSet {
	names := make(fileSet, len(files))
	rootFinder := newModuleRootFinder()

	for _, file := range files {
		if file == nil {
			continue
		}

		if tokFile := fset.File(file.Pos()); tokFile != nil {
			name := filepath.Clean(tokFile.Name())
			if matcher != nil && matcher.matchesFile(name, rootFinder) {
				continue
			}

			names[name] = struct{}{}
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
