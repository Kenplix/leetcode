package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx, tail := 0, m
next:
	for _, num2 := range nums2 {
		for {
			switch num1 := nums1[idx]; {
			case num2 < num1:
				// shift items
				copy(nums1[idx+1:], nums1[idx:])
				fallthrough
			case idx == tail:
				nums1[idx] = num2
				tail++
				idx++
				continue next
			default:
				// find the first other value
				for i, num := range nums1[idx+1:] {
					if num != num1 {
						idx += i
						break
					}
				}
			}

			idx++
		}
	}
}
