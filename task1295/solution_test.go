package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "first",
			nums:     []int{12, 345, 2, 6, 7896},
			expected: 2,
		},
		{
			name:     "second",
			nums:     []int{555, 901, 482, 1771},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findNumbers(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkFindNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findNumbers(testcaseGenerator(maxNumsLen, maxNumsLen, minNumValue, maxNumValue))
	}
}
