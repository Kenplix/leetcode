package main

import (
	"fmt"
	"leetcode/helpers"
	"math/rand"
	"sort"
	"time"
)

// Constraints:
const (
	minNumsLen = 1
	maxNumsLen = 10e4

	minNumValue = -10e4
	maxNumValue = +10e4
)

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	output := make([]int, len(nums))
	last := len(output) - 1
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			output[last] = leftSquare
			left++
		} else if leftSquare <= rightSquare {
			output[last] = rightSquare
			right--
		}
		last--
	}

	return output
}

func genTestCase(minNumsLen, maxNumsLen, minNumValue, maxNumValue int) []int {
	nums := make([]int, helpers.RandIntInRange(minNumsLen, maxNumsLen))
	for idx := 0; idx < len(nums); idx++ {
		nums[idx] = helpers.RandIntInRange(minNumValue, maxNumValue)
	}
	sort.Ints(nums)

	return nums
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("result:", sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println("result:", sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println("result:", sortedSquares(genTestCase(10, 10, -10, 10)))
}
