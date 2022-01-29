package main

// removeElement faster than 100% of Golang submissions
func removeElement(nums []int, val int) int {
	leftPtr, rightPtr := 0, len(nums)-1
	for leftPtr <= rightPtr {
		if leftNum := nums[leftPtr]; leftNum == val {
			if rightNum := nums[rightPtr]; rightNum != val {
				nums[leftPtr] = rightNum
				leftPtr++
			}

			nums = nums[:rightPtr]
			rightPtr--
			continue
		}

		leftPtr++
	}

	return len(nums)
}
