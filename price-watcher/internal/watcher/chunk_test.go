package watcher

import (
	"reflect"
	"testing"
)

func TestChunkSymbols(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		chunkSize int
		expected  [][]string
	}{
		{
			name:      "even split",
			input:     []string{"a", "b", "c", "d"},
			chunkSize: 2,
			expected:  [][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			name:      "uneven split",
			input:     []string{"a", "b", "c", "d", "e"},
			chunkSize: 2,
			expected:  [][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			name:      "chunk size > input length",
			input:     []string{"a", "b"},
			chunkSize: 10,
			expected:  [][]string{{"a", "b"}},
		},
		{
			name:      "chunk size == input length",
			input:     []string{"a", "b", "c"},
			chunkSize: 3,
			expected:  [][]string{{"a", "b", "c"}},
		},
		{
			name:      "chunk size is 1",
			input:     []string{"a", "b"},
			chunkSize: 1,
			expected:  [][]string{{"a"}, {"b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := chunkSymbols(tt.input, tt.chunkSize)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
