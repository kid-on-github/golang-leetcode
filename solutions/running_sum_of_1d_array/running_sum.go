package runningsum

func runningSum(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	summation := []int{nums[0]}
	for _, num := range nums[1:] {
		prev := summation[len(summation)-1]
		summation = append(summation, prev+num)
	}

	return summation
}
