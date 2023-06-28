package findpivotindex

func pivotIndex(nums []int) int {
	runningSum := []int{nums[0]}

	for i, num := range nums[1:] {
		runningSum = append(runningSum, num+runningSum[i])
	}

	total := runningSum[len(runningSum)-1]

	pivot := total

	for i, num := range runningSum {
		if num == pivot {
			return i
		}

		pivot = total - num
	}

	return -1
}
