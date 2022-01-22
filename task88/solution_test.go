package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1, 2, 3, 4, 5, 6, 0},
			m:        6,
			nums2:    []int{3},
			n:        1,
			expected: []int{1, 2, 3, 3, 4, 5, 6},
		},
		{
			nums1:    []int{-1, 0, 1, 2, 3, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{-1, 0, 1, 4},
			n:        4,
			expected: []int{-1, -1, 0, 0, 1, 1, 2, 3, 4},
		},
		{
			nums1:    []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{-1, -1, 0, 0, 1, 2, 4, 5},
			n:        8,
			expected: []int{-1, -1, -1, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5},
		},
		{
			nums1:    []int{1, 2, 0, 0, 0},
			m:        2,
			nums2:    []int{3, 4, 5},
			n:        3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			assert.Equal(t, tc.expected, tc.nums1)
		})
	}
}
