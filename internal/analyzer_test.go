package camellia_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	camelliainternal "github.com/caelaxie/camellia/internal"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, camelliainternal.Analyzer, "a")
}
