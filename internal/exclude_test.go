package camellia

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizeExcludePattern(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "trim and clean relative path",
			input: " ./internal/testdata/../fixtures/*.go ",
			want:  "internal/fixtures/*.go",
		},
		{
			name:    "reject empty pattern",
			input:   "   ",
			wantErr: true,
		},
		{
			name:    "reject parent escape",
			input:   "../fixtures",
			wantErr: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := normalizeExcludePattern(test.input)
			if test.wantErr {
				if err == nil {
					t.Fatalf("normalizeExcludePattern(%q) error = nil, want error", test.input)
				}

				return
			}

			if err != nil {
				t.Fatalf("normalizeExcludePattern(%q) error = %v", test.input, err)
			}

			if got != test.want {
				t.Fatalf("normalizeExcludePattern(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestNormalizeExcludePatternRejectsAbsolutePath(t *testing.T) {
	t.Parallel()

	absolutePath := filepath.Join(string(filepath.Separator), "tmp", "excluded.go")

	_, err := normalizeExcludePattern(absolutePath)
	if err == nil {
		t.Fatal("normalizeExcludePattern() error = nil, want error")
	}
}

func TestNewExcludeMatcher(t *testing.T) {
	t.Parallel()

	t.Run("nil for empty patterns", func(t *testing.T) {
		t.Parallel()

		matcher, err := newExcludeMatcher(nil)
		if err != nil {
			t.Fatalf("newExcludeMatcher(nil) error = %v", err)
		}

		if matcher != nil {
			t.Fatalf("newExcludeMatcher(nil) = %#v, want nil", matcher)
		}
	})

	t.Run("reject malformed glob", func(t *testing.T) {
		t.Parallel()

		_, err := newExcludeMatcher([]string{"["})
		if err == nil {
			t.Fatal("newExcludeMatcher() error = nil, want malformed pattern error")
		}
	})
}

func TestExcludePatternMatches(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		pattern  excludePattern
		path     string
		expected bool
	}{
		{
			name: "literal file exact match",
			pattern: excludePattern{
				pattern: "internal/testdata/file.go",
				literal: true,
			},
			path:     "internal/testdata/file.go",
			expected: true,
		},
		{
			name: "literal file does not match sibling",
			pattern: excludePattern{
				pattern: "internal/testdata/file.go",
				literal: true,
			},
			path:     "internal/testdata/other.go",
			expected: false,
		},
		{
			name: "literal directory matches subtree",
			pattern: excludePattern{
				pattern: "internal/testdata",
				literal: true,
			},
			path:     "internal/testdata/fixtures/file.go",
			expected: true,
		},
		{
			name: "glob matches subtree",
			pattern: excludePattern{
				pattern: "internal/**/excluded.go",
				literal: false,
			},
			path:     "internal/testdata/excluded.go",
			expected: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.pattern.matches(test.path)
			if got != test.expected {
				t.Fatalf("excludePattern.matches(%q) = %v, want %v", test.path, got, test.expected)
			}
		})
	}
}

func TestModuleRootFinderRelativePath(t *testing.T) {
	t.Parallel()

	rootDir := t.TempDir()
	nestedDir := filepath.Join(rootDir, "internal", "nested")
	filePath := filepath.Join(nestedDir, "sample.go")

	if err := os.MkdirAll(nestedDir, 0o750); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}

	if err := os.WriteFile(filepath.Join(rootDir, "go.mod"), []byte("module example.com/test\n\ngo 1.25.4\n"), 0o600); err != nil {
		t.Fatalf("WriteFile(go.mod) error = %v", err)
	}

	if err := os.WriteFile(filePath, []byte("package nested\n"), 0o600); err != nil {
		t.Fatalf("WriteFile(sample.go) error = %v", err)
	}

	rootFinder := newModuleRootFinder()

	relativePath, ok := rootFinder.relativePath(filePath)
	if !ok {
		t.Fatal("relativePath() ok = false, want true")
	}

	if relativePath != "internal/nested/sample.go" {
		t.Fatalf("relativePath() = %q, want %q", relativePath, "internal/nested/sample.go")
	}
}
