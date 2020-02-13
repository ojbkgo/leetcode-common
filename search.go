package common

func BinarySearchInt(data []int, target int) (int, bool) {

	if len(data) == 0 {
		return -1, false
	}

	var l, r, m int

	l = 0
	r = len(data) - 1
	m = (l + r) / 2

	for l < r {
		if target > data[m] {
			l = m+1
		} else if target < data[m] {
			r = m
		} else {
			return m, true
		}
		m = (l + r) / 2
	}

	return l, false
}

func BinarySearchString(data []int, target int) (int, bool) {
	return -1, false
}
