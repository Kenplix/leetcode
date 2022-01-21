package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{},
			expected: []int{},
		},
		{
			nums:     []int{1},
			expected: []int{1},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{0, 1},
			expected: []int{0, 0},
		},
		{
			nums:     []int{1, 0},
			expected: []int{1, 0},
		},
		{
			nums:     []int{0, 1, 0},
			expected: []int{0, 0, 1},
		},
		{
			nums:     []int{0, 1, 2},
			expected: []int{0, 0, 1},
		},
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			nums:     []int{1, 0, 2, 3, 0, 4, 5, 0, 6, 7},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4, 5, 0},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			duplicateZeros(tc.nums)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}

func BenchmarkDuplicateZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		duplicateZeros(genTestCase(maxNumsLen, maxNumsLen, minNumValue, maxNumValue))
	}
}
