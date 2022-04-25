package first_bad_version

var bad_version int = 0

func isBadVersion(version int) bool {
	if version >= bad_version {
		return true
	}

	return false
}

func firstBadVersion(n int) int {
	low := 0
	high := n
	good := 0

	for low <= high {
		middle := (low + high) / 2
		result := isBadVersion(middle)

		if result {
			if middle-1 == good {
				return middle
			}
			high = middle
		} else {
			good = middle
			low = middle + 1
		}
	}

	return 0
}
