package camellia_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/caelaxie/camellia/internal/camellia"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, camellia.Analyzer, "a")
}
