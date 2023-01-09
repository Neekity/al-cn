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

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lenH, lenN := len(haystack), len(needle)
	var nextArr [50001]int
	nextArr[0] = -1
	if lenN > 1 {
		nextArr[1] = 0
		pos, cn := 2, 0
		for pos < lenN {
			if needle[pos-1] == needle[cn] {
				cn++
				nextArr[pos] = cn
				pos++
			} else if cn > 0 {
				cn = nextArr[cn]
			} else {
				nextArr[pos] = 0
				pos++
			}
		}
	}
	hi, ni := 0, 0
	for hi < lenH && ni < lenN {
		if haystack[hi] == needle[ni] {
			ni++
			hi++
		} else if nextArr[ni] == -1 {
			hi++
		} else {
			ni = nextArr[ni]
		}
	}
	if ni == lenN {
		return hi - ni
	}
	return -1
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		return float64(findKthNumber(nums1, nums2, totalLength/2+1))
	} else {
		res = float64(findKthNumber(nums1, nums2, totalLength/2+1)+findKthNumber(nums1, nums2, totalLength/2)) / 2
	}
	return res
}

func findKthNumber(nums1 []int, nums2 []int, kth int) int {
	lenShort, lenLong := len(nums1), len(nums2)
	if lenShort > lenLong {
		return findKthNumber(nums2, nums1, kth)
	}
	if lenShort == 0 {
		return nums2[kth-1]
	}

	if kth <= lenShort {
		return findMidNumber(nums1, nums2, 0, 0, kth, kth)
	}

	if kth > lenLong {
		if nums2[kth-lenShort-1] >= nums1[lenShort-1] {
			return nums2[kth-lenShort-1]
		}

		if nums1[kth-lenLong-1] >= nums2[lenLong-1] {
			return nums1[kth-lenLong-1]
		}
		return findMidNumber(nums1, nums2, kth-lenLong, kth-lenShort, lenShort, lenLong)
	}

	if nums1[lenShort-1] <= nums2[kth-lenShort-1] {
		return nums2[kth-lenShort-1]
	}

	return findMidNumber(nums1, nums2, 0, kth-lenShort, lenShort, kth)
}

func findMidNumber(nums1 []int, nums2 []int, l1 int, l2 int, r1 int, r2 int) int {
	mid1, mid2, offset := 0, 0, 0

	for {
		if (r1-l1)%2 == 0 {
			mid1, mid2 = (r1+l1)/2-1, (r2+l2)/2-1
		} else {
			mid1, mid2 = (r1+l1)/2, (r2+l2)/2
		}
		if l1 >= r1-1 {
			break
		}
		offset = 1 - (r1-l1)%2

		if nums1[mid1] < nums2[mid2] {
			l1 = mid1 + offset
			r2 = mid2 + 1
		} else if nums1[mid1] > nums2[mid2] {
			r1 = mid1 + 1
			l2 = mid2 + offset
		} else {
			return nums1[mid1]
		}
	}

	return common.Min(nums1[l1], nums2[l2])
}
