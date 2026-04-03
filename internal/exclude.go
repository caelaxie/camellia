package camellia

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	doublestar "github.com/bmatcuk/doublestar/v4"
)

type excludeMatcher struct {
	patterns []excludePattern
}

type excludePattern struct {
	pattern string
	literal bool
}

func newExcludeMatcher(rawPatterns []string) (*excludeMatcher, error) {
	if len(rawPatterns) == 0 {
		return nil, nil
	}

	patterns := make([]excludePattern, 0, len(rawPatterns))
	for _, rawPattern := range rawPatterns {
		pattern, err := normalizeExcludePattern(rawPattern)
		if err != nil {
			return nil, err
		}

		literal := !hasGlobMeta(pattern)
		if !literal && !doublestar.ValidatePattern(pattern) {
			return nil, fmt.Errorf("invalid exclude pattern %q", rawPattern)
		}

		patterns = append(patterns, excludePattern{
			pattern: pattern,
			literal: literal,
		})
	}

	return &excludeMatcher{patterns: patterns}, nil
}

func normalizeExcludePattern(rawPattern string) (string, error) {
	pattern := strings.TrimSpace(rawPattern)
	if pattern == "" {
		return "", fmt.Errorf("exclude pattern must not be empty")
	}

	if filepath.IsAbs(pattern) {
		return "", fmt.Errorf("exclude pattern %q must be relative to module root", rawPattern)
	}

	pattern = filepath.ToSlash(pattern)
	pattern = path.Clean(pattern)

	if pattern == "." || pattern == ".." || strings.HasPrefix(pattern, "../") {
		return "", fmt.Errorf("exclude pattern %q must stay within module root", rawPattern)
	}

	return pattern, nil
}

func hasGlobMeta(pattern string) bool {
	return strings.ContainsAny(pattern, "*?[{")
}

func (m *excludeMatcher) matchesFile(fileName string, rootFinder *moduleRootFinder) bool {
	if m == nil {
		return false
	}

	relativePath, ok := rootFinder.relativePath(fileName)
	if !ok {
		return false
	}

	for _, pattern := range m.patterns {
		if pattern.matches(relativePath) {
			return true
		}
	}

	return false
}

func (p excludePattern) matches(relativePath string) bool {
	if p.literal {
		return relativePath == p.pattern || strings.HasPrefix(relativePath, p.pattern+"/")
	}

	return doublestar.MatchUnvalidated(p.pattern, relativePath)
}

type moduleRootFinder struct {
	roots   map[string]string
	missing map[string]struct{}
}

func newModuleRootFinder() *moduleRootFinder {
	return &moduleRootFinder{
		roots:   make(map[string]string),
		missing: make(map[string]struct{}),
	}
}

func (f *moduleRootFinder) relativePath(fileName string) (string, bool) {
	absolutePath, err := filepath.Abs(fileName)
	if err != nil {
		return "", false
	}

	root, ok := f.findRoot(filepath.Dir(absolutePath))
	if !ok {
		return "", false
	}

	relativePath, err := filepath.Rel(root, absolutePath)
	if err != nil {
		return "", false
	}

	relativePath = filepath.ToSlash(relativePath)
	if relativePath == ".." || strings.HasPrefix(relativePath, "../") {
		return "", false
	}

	return relativePath, true
}

func (f *moduleRootFinder) findRoot(dir string) (string, bool) {
	searchedDirs := make([]string, 0, 4)

	for {
		if root, ok := f.roots[dir]; ok {
			for _, searchedDir := range searchedDirs {
				f.roots[searchedDir] = root
			}

			return root, true
		}

		if _, ok := f.missing[dir]; ok {
			for _, searchedDir := range searchedDirs {
				f.missing[searchedDir] = struct{}{}
			}

			return "", false
		}

		searchedDirs = append(searchedDirs, dir)

		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			for _, searchedDir := range searchedDirs {
				f.roots[searchedDir] = dir
			}

			return dir, true
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			for _, searchedDir := range searchedDirs {
				f.missing[searchedDir] = struct{}{}
			}

			return "", false
		}

		dir = parentDir
	}
}
