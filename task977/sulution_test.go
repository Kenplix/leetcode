package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{},
			expected: []int{},
		},
		{
			nums:     []int{-7},
			expected: []int{49},
		},
		{
			nums:     []int{-7, 11},
			expected: []int{49, 121},
		},
		{
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := sortedSquares(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkSortedSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedSquares(genTestCase(maxNumsLen, maxNumsLen, minNumValue, maxNumValue))
	}
}
