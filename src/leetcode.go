package src

func TwoSum(nums []int, target int) []int {
	var searchArray map[int]int
	searchArray = make(map[int]int)
	for idx, num := range nums {
		_, ok := searchArray[target-num]
		if ok == true {
			return []int{searchArray[target-num], idx}
		}
		searchArray[target-num] = idx
	}
	return []int{0, 0}
}
