// go test -v
package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1: reachable",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2: stuck at zero",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Cannot move from start",
			nums:     []int{0, 1},
			expected: false,
		},
		{
			name:     "Step-by-step path",
			nums:     []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Big jump at start",
			nums:     []int{5, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Stuck before last",
			nums:     []int{2, 0, 0, 1},
			expected: false,
		},
		{
			name:     "Exact jump to end",
			nums:     []int{3, 0, 0, 0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
