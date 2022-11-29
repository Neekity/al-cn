package src

import "fmt"

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

func BinarySearch(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func MaxSubArray(nums []int) int {
	res := nums[0]
	n := len(nums)
	for cur := 1; cur < n; cur++ {
		if nums[cur]+nums[cur-1] > nums[cur] {
			nums[cur] += nums[cur-1]
		}
		if nums[cur] > res {
			res = nums[cur]
		}
	}
	fmt.Println(nums)
	return res
}
