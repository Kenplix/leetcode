package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "first",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "second",
			nums:     []int{-7},
			expected: []int{49},
		},
		{
			name:     "third",
			nums:     []int{-7, 11},
			expected: []int{49, 121},
		},
		{
			name:     "fourth",
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			name:     "fifth",
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortedSquares(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkFindNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedSquares(testcaseGenerator(maxNumsLen, maxNumsLen, minNumValue, maxNumValue))
	}
}
