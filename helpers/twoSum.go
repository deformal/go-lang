package helpers

func TwoSum(nums []int, target int) []int {
	answer := []int{}
	iMap := make(map[int]int)
	for i, v := range nums {
		//ANCHOR basically it checks if the two numbers
		//which can make the target sum
		//are availabe in the new iMap or not
		if mi, ok := iMap[target-v]; ok {
			answer = append(answer, mi, i)
		}
		//ANCHOR - stores the numbers
		//as key and the indexes as their value
		iMap[v] = i
	}

	return answer
}
