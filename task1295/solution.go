package main

import (
	"fmt"
	"math/rand"
)

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

// randIntInRange returns number in range [from, to]
func randIntInRange(min, max int) int {
	if min > max {
		panic("max value should be greater than min")
	}

	return min + rand.Intn(max-min+1)
}

func testcaseGenerator(minNumsLen, maxNumsLen, minNumValue, maxNumValue int) []int {
	nums := make([]int, randIntInRange(minNumsLen, maxNumsLen))

	for idx := 0; idx < len(nums); idx++ {
		nums[idx] = randIntInRange(minNumValue, maxNumValue)
	}

	return nums
}

func main() {
	fmt.Println("result:", findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println("result:", findNumbers([]int{555, 901, 482, 1771}))
}
