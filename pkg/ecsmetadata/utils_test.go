package ecsmetadata

import "testing"

func TestParsePathNames(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "normal case with multiple paths",
			input: "path1\npath2/\n path3 \n",
			expected: []string{
				"path1",
				"path2",
				"path3",
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "input with only whitespace",
			input:    "   \n\t\n   ",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parsePathNames(tt.input)
			if !equalStringSlices(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseJSONStringArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []string
		hasError bool
	}{
		{
			name:     "normal case",
			input:    []byte(`["item1", "item2", "item3"]`),
			expected: []string{"item1", "item2", "item3"},
			hasError: false,
		},
		{
			name:     "empty JSON array",
			input:    []byte(`[]`),
			expected: []string{},
			hasError: false,
		},
		{
			name:     "invalid JSON",
			input:    []byte(`[invalid-json`),
			expected: nil,
			hasError: true,
		},
		{
			name:     "non-array JSON",
			input:    []byte(`"not-an-array"`),
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseJSONStringArray(tt.input)
			if (err != nil) != tt.hasError {
				t.Errorf("expected error: %v, got: %v", tt.hasError, err)
			}
			if !equalStringSlices(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper function to compare two string slices
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
