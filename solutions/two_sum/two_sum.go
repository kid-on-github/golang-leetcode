package twosum

func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)

	for i, num := range nums {
		desired := target - num

		if desiredIndex, desiredInMap := numIndex[desired]; desiredInMap {
			return []int{desiredIndex, i}
		}

		numIndex[num] = i
	}

	return []int{-1, -1}
}
