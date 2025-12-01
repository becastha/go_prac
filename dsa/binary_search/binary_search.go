package binary_search

func Binary_search(sortedList []int, num int) int {
	low := 0
	high := len(sortedList) - 1

	// find the mid
	for low <= high {
		mid := low + (high-low)/2
		midNum := sortedList[mid]

		if midNum == num {
			return mid
		} else if midNum > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
