package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	r := [][]int{}
	if len(intervals) == 0 {
		r = append(r, newInterval)
		return r
	}

	left, right := intervals[0][0], intervals[0][1]

	if left > newInterval[1] {
		r = append(r, []int{newInterval[0], newInterval[1]})
	} else if left > newInterval[0] {
		left = newInterval[0]
	}
	for i := 0; i < len(intervals); i++ {
		nextLeft, nextRight := 1000000, 1000000
		if i+1 < len(intervals) {
			nextLeft, nextRight = intervals[i+1][0], intervals[i+1][1]
		}

		if left > newInterval[1] {
			if len(r) == 0 || r[len(r)-1][1] < newInterval[0] {
				r = append(r, []int{newInterval[0], newInterval[1]})
			}
			r = append(r, []int{left, right})
			left, right = nextLeft, nextRight
		} else if left >= newInterval[0] && right <= newInterval[1] {
			left, right = nextLeft, nextRight
		} else if left >= newInterval[0] && right >= newInterval[1] {
			r = append(r, []int{newInterval[0], right})
			left, right = nextLeft, nextRight
		} else if left < newInterval[0] && right > newInterval[1] {
			r = append(r, []int{left, right})
			left, right = nextLeft, nextRight
		} else if newInterval[0] <= right && newInterval[1] >= right {
			newInterval[0] = left
			left, right = nextLeft, nextRight
		} else {
			r = append(r, []int{left, right})
			left, right = nextLeft, nextRight
		}
	}
	if left < 1000000 {
		if left > newInterval[1] {
			if r[len(r)-1][1] < newInterval[0] {
				r = append(r, []int{newInterval[0], newInterval[1]})
			}
			r = append(r, []int{left, right})
		} else if left >= newInterval[0] && right <= newInterval[1] {
			r = append(r, []int{newInterval[0], newInterval[1]})
		} else if left >= newInterval[0] && right >= newInterval[1] {
			r = append(r, []int{newInterval[0], right})
		} else if left < newInterval[0] && right > newInterval[1] {
			r = append(r, []int{left, right})
		} else if newInterval[0] <= right && newInterval[1] >= right {
			r = append(r, []int{left, newInterval[1]})
		} else {
			r = append(r, []int{left, right})
			r = append(r, []int{newInterval[0], newInterval[1]})
		}
	}
	if len(r) == 0 || r[len(r)-1][1] < newInterval[0] {
		r = append(r, []int{newInterval[0], newInterval[1]})
	}

	return r
}
