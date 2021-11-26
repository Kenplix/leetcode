package main

import (
	"fmt"
	"leetcode/helpers"
	"math/rand"
	"time"
)

// Constraints:
const (
	minNumsLen = 1
	maxNumsLen = 500

	minNumValue = 1
	maxNumValue = 10e5
)

func findNumbers(nums []int) int {
	res, rank, num := 0, uint8(0), 0
	for idx := 0; idx < len(nums); idx++ {
		num, rank = nums[idx], 0
		for num != 0 {
			num /= 10
			rank++
		}

		if rank%2 == 0 {
			res++
		}
	}

	return res
}

func testcaseGenerator(minNumsLen, maxNumsLen, minNumValue, maxNumValue int) []int {
	nums := make([]int, helpers.RandIntInRange(minNumsLen, maxNumsLen))
	for idx := 0; idx < len(nums); idx++ {
		nums[idx] = helpers.RandIntInRange(minNumValue, maxNumValue)
	}

	return nums
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("result:", findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println("result:", findNumbers([]int{555, 901, 482, 1771}))
	fmt.Println("result:", findNumbers(testcaseGenerator(10, 10, minNumValue, maxNumValue)))
}
