package camellia_test

import (
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	camelliainternal "github.com/caelaxie/camellia/internal"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, camelliainternal.Analyzer, "a")
}

func TestAnalyzerExcludeFile(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analyzer := newAnalyzerForTest(t, camelliainternal.Config{
		Exclude: []string{"internal/testdata/src/excludefile/excluded.go"},
	})

	analysistest.Run(t, testdata, analyzer, "excludefile")
}

func TestAnalyzerExcludeDirectory(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analyzer := newAnalyzerForTest(t, camelliainternal.Config{
		Exclude: []string{"internal/testdata/src/excludeddir"},
	})

	analysistest.Run(t, testdata, analyzer, "excludeddir")
}

func TestAnalyzerExcludeDirectoryGlob(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analyzer := newAnalyzerForTest(t, camelliainternal.Config{
		Exclude: []string{"internal/testdata/src/excludeddir/**"},
	})

	analysistest.Run(t, testdata, analyzer, "excludeddir")
}

func newAnalyzerForTest(t *testing.T, config camelliainternal.Config) *analysis.Analyzer {
	t.Helper()

	analyzer, err := camelliainternal.NewAnalyzer(config)
	if err != nil {
		t.Fatalf("NewAnalyzer() error = %v", err)
	}

	return analyzer
}
