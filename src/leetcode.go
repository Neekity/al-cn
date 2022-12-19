package src

import (
	"fmt"
	"math"
	"neekity.com/al-cn/src/common"
)

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

func PlusOne(digits []int) []int {
	n := len(digits)
	flag := 1
	for i := n - 1; i >= 0; i-- {
		tmp := digits[i] + flag
		flag = tmp / 10
		digits[i] = tmp % 10
		if flag == 0 {
			break
		}
	}
	if flag == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	curr := head
	var prev *common.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func MergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	dump := &common.ListNode{-1, nil}
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur, list1 = list1, list1.Next
			cur.Next = nil
		} else if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = list2, list2.Next
			cur.Next = nil
		}
	}
	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return dump.Next
}

func Reverse(x int32) int {
	var res int32
	var digit int32
	res = 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit = x % 10
		x = x / 10
		res = 10*res + digit
	}
	return int(res)
}

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	preOne := 1
	preTwo := 2
	for i := 2; i < n; i++ {
		preOne, preTwo = preTwo, preOne+preTwo
	}
	return preTwo
}
