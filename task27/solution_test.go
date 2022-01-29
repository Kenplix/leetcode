package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}

	testCases := []struct {
		name   string
		input  args
		output int
	}{
		{
			input: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			output: 2,
		},
		{
			input: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			output: 5,
		},
		{
			input: args{
				nums: []int{0, 1, 2, 2, 2, 0, 4, 2},
				val:  2,
			},
			output: 4,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := removeElement(tc.input.nums, tc.input.val)
			assert.Equal(t, tc.output, actual)
		})
	}
}
