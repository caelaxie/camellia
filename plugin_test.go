package camellia

import (
	"strings"
	"testing"
)

func TestNewAcceptsExcludeSettings(t *testing.T) {
	t.Parallel()

	plugin, err := New(map[string]any{
		"exclude": []any{"internal/testdata", "generated/**/*.go"},
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	analyzers, err := plugin.BuildAnalyzers()
	if err != nil {
		t.Fatalf("BuildAnalyzers() error = %v", err)
	}

	if len(analyzers) != 1 {
		t.Fatalf("BuildAnalyzers() len = %d, want 1", len(analyzers))
	}
}

func TestNewRejectsMalformedExcludePattern(t *testing.T) {
	t.Parallel()

	_, err := New(map[string]any{
		"exclude": []string{"["},
	})
	if err == nil {
		t.Fatal("New() error = nil, want malformed pattern error")
	}

	if !strings.Contains(err.Error(), `invalid exclude pattern "["`) {
		t.Fatalf("New() error = %q, want malformed pattern message", err)
	}
}
