package main

import (
	"leetcode/helpers"
	"math/rand"
	"time"
)

// Constraints:
const (
	minNumsLen = 1
	maxNumsLen = 10e4

	minNumValue = 0
	maxNumValue = 9
)

func duplicateZeros(arr []int) {
	shifts, last := 0, len(arr)-1
	for idx, val := range arr {
		if idx == last {
			if val == 0 {
				last--
			}
			break
		}

		if val == 0 {
			arr[last] = -1
			last--
			shifts++
		}

	}
	//fmt.Println(arr)

	border := len(arr)
	for ; shifts > 0; last-- {
		if arr[last] == 0 {
			shift(arr[:border], last, shifts)
			border = last + shifts
			shifts--
		}
	}

	//fmt.Println("result:", arr)
}

func shift(arr []int, pos, times int) {
	//fmt.Printf("pos: %d, times: %d\n", pos, times)
	for idx := len(arr) - 1; idx >= pos+times; idx-- {
		arr[idx] = arr[idx-times]
	}
	//fmt.Println(arr)
}

func genTestCase(minNumsLen, maxNumsLen, minNumValue, maxNumValue int) []int {
	nums := make([]int, helpers.RandIntInRange(minNumsLen, maxNumsLen))
	for idx := 0; idx < len(nums); idx++ {
		nums[idx] = helpers.RandIntInRange(minNumValue, maxNumValue)
	}

	return nums
}

func main() {
	rand.Seed(time.Now().UnixNano())
	duplicateZeros([]int{1, 2, 3})
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
	duplicateZeros(genTestCase(minNumsLen, maxNumsLen, minNumValue, maxNumValue))
}
