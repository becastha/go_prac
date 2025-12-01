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

func BinarySearchRecursive(sortedList []int, num int) int {
	return binarySearchHelper(sortedList, num, 0, len(sortedList)-1)
}

func binarySearchHelper(arr []int, num, low, high int) int {
	if low > high {
		return -1
	}

	// find the middle index
	mid := low + (high-low)/2

	// if the middle element matches the numerb
	if arr[mid] == num {
		return mid
	}

	// if the middle element is greater that the target
	if arr[mid] > num {
		return binarySearchHelper(arr, num, low, mid-1)
	}

	return binarySearchHelper(arr, num, mid+1, high)
}
