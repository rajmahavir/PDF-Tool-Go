package pdf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParsePageNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "Single page",
			input:    "5",
			expected: []int{5},
		},
		{
			name:     "Multiple pages",
			input:    "1,3,5",
			expected: []int{1, 3, 5},
		},
		{
			name:     "Pages with spaces",
			input:    "1, 3, 5",
			expected: []int{1, 3, 5},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []int{},
		},
		{
			name:     "Invalid numbers ignored",
			input:    "1,abc,3",
			expected: []int{1, 3},
		},
		{
			name:     "Zero and negative ignored",
			input:    "0,-1,2,3",
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParsePageNumbers(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ParsePageNumbers(%q) = %v, want %v", tt.input, result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("ParsePageNumbers(%q) = %v, want %v", tt.input, result, tt.expected)
					return
				}
			}
		})
	}
}

func TestGetPagesToKeep(t *testing.T) {
	tests := []struct {
		name          string
		totalPages    int
		pagesToRemove []int
		expected      []int
	}{
		{
			name:          "Remove single page from middle",
			totalPages:    5,
			pagesToRemove: []int{3},
			expected:      []int{1, 2, 4, 5},
		},
		{
			name:          "Remove multiple pages",
			totalPages:    5,
			pagesToRemove: []int{2, 4},
			expected:      []int{1, 3, 5},
		},
		{
			name:          "Remove first page",
			totalPages:    3,
			pagesToRemove: []int{1},
			expected:      []int{2, 3},
		},
		{
			name:          "Remove last page",
			totalPages:    3,
			pagesToRemove: []int{3},
			expected:      []int{1, 2},
		},
		{
			name:          "No pages to remove",
			totalPages:    3,
			pagesToRemove: []int{},
			expected:      []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPagesToKeep(tt.totalPages, tt.pagesToRemove)
			if len(result) != len(tt.expected) {
				t.Errorf("GetPagesToKeep(%d, %v) = %v, want %v", tt.totalPages, tt.pagesToRemove, result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("GetPagesToKeep(%d, %v) = %v, want %v", tt.totalPages, tt.pagesToRemove, result, tt.expected)
					return
				}
			}
		})
	}
}

func TestCreatePageRanges(t *testing.T) {
	tests := []struct {
		name     string
		pages    []int
		expected []string
	}{
		{
			name:     "Single page",
			pages:    []int{5},
			expected: []string{"5"},
		},
		{
			name:     "Consecutive pages",
			pages:    []int{1, 2, 3, 4, 5},
			expected: []string{"1-5"},
		},
		{
			name:     "Non-consecutive pages",
			pages:    []int{1, 3, 5},
			expected: []string{"1", "3", "5"},
		},
		{
			name:     "Mixed consecutive and non-consecutive",
			pages:    []int{1, 2, 3, 5, 7, 8, 9},
			expected: []string{"1-3", "5", "7-9"},
		},
		{
			name:     "Two consecutive pages",
			pages:    []int{1, 2},
			expected: []string{"1-2"},
		},
		{
			name:     "Empty pages",
			pages:    []int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreatePageRanges(tt.pages)
			if len(result) != len(tt.expected) {
				t.Errorf("CreatePageRanges(%v) = %v, want %v", tt.pages, result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("CreatePageRanges(%v) = %v, want %v", tt.pages, result, tt.expected)
					return
				}
			}
		})
	}
}

func TestSaveFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test-savefile-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testData := "test content"
	reader := strings.NewReader(testData)
	filePath := filepath.Join(tempDir, "test.txt")

	err = SaveFile(reader, filePath)
	if err != nil {
		t.Errorf("SaveFile() error = %v", err)
		return
	}

	// Verify file was created
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("SaveFile() did not create file at %s", filePath)
		return
	}

	// Verify file contents
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read saved file: %v", err)
		return
	}

	if string(content) != testData {
		t.Errorf("SaveFile() saved content = %q, want %q", string(content), testData)
	}
}

func TestSaveFileInvalidPath(t *testing.T) {
	reader := strings.NewReader("test")
	err := SaveFile(reader, "/invalid/path/that/does/not/exist/file.txt")
	if err == nil {
		t.Error("SaveFile() expected error for invalid path, got nil")
	}
}

// Integration test helper - only run if test PDFs are available
func skipIfNoPDFTestData(t *testing.T) {
	if _, err := os.Stat("testdata"); os.IsNotExist(err) {
		t.Skip("Skipping test: testdata directory not found")
	}
}
