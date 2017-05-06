package main

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	if high == 0 {
		return false
	}

	for low <= high {
		mid := low + (low-high)/2
		if nums[mid] == target {
			return true
		} else {
			if nums[low] == nums[mid] {
				low += 1
			} else if nums[low] < nums[mid] {
				if nums[low] <= target && target < nums[mid] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			} else {
				if nums[mid] < target && target <= nums[high] {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		}
	}

	return false
}

func main() {
	// grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	in := []int{1}
}

func max(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
