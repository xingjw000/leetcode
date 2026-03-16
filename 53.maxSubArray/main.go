package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	// 遍历
	// l := len(nums)
	// max_sum := -math.MaxInt
	// for i := 0; i < l; i++ {
	// 	sum := 0
	// 	for j := i; j >= 0; j-- {
	// 		sum += nums[j]
	// 		if max_sum < sum {
	// 			max_sum = sum
	// 		}
	// 	}
	// }
	// return max_sum

	// 贪心算法
	// right := 0
	// l := len(nums)
	// max_sum := -math.MaxInt
	// sum := 0
	// for right < l {
	// 	sum += nums[right]
	// 	right++
	// 	if max_sum < sum {
	// 		max_sum = sum
	// 	}
	// 	if sum <= 0 {
	// 		sum = 0
	// 	}
	// }
	// return max_sum

	return merge(nums, 0, len(nums)-1).maxSum

}

// func merge_sum(nums []int, l int, r int) int {
// 	if l == r {
// 		return nums[l]
// 	}
// 	mid := (l + r) / 2
// 	l_max_sum := merge_sum(nums, l, mid)
// 	r_max_sum := merge_sum(nums, mid+1, r)
// 	m_max_sum := cross_left_right_max_sum(nums, l, r)

// 	max_sum := l_max_sum

// 	if max_sum < r_max_sum {
// 		max_sum = r_max_sum
// 	}

// 	if max_sum < m_max_sum {
// 		max_sum = m_max_sum
// 	}

// 	return max_sum
// }

// func cross_left_right_max_sum(nums []int, l int, r int) int {
// 	mid := (l + r) / 2
// 	r_max_sum := nums[mid+1]
// 	sum := 0
// 	for i := mid + 1; i <= r; i++ {
// 		sum += nums[i]
// 		if r_max_sum < sum {
// 			r_max_sum = sum
// 		}
// 	}

// 	sum = 0
// 	l_max_sum := nums[l]
// 	for i := mid; i >= l; i-- {
// 		sum += nums[i]
// 		if l_max_sum < sum {
// 			l_max_sum = sum
// 		}
// 	}

// 	return l_max_sum + r_max_sum
// }

func merge(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}

	mid := (l + r) / 2
	lStatus := merge(nums, l, mid)
	rStatus := merge(nums, mid+1, r)
	return pushUp(lStatus, rStatus)
}

func pushUp(lStatus, rStatus Status) Status {
	lSum := myMax(lStatus.lSum, lStatus.sum+rStatus.lSum)
	rSum := myMax(lStatus.rSum+rStatus.sum, rStatus.rSum)
	maxSum := myMax(lStatus.rSum+rStatus.lSum, myMax(lStatus.maxSum, rStatus.maxSum))
	sum := lStatus.sum + rStatus.sum
	return Status{lSum, rSum, maxSum, sum}
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Status struct {
	lSum, rSum, maxSum, sum int
}

// func maxSubArray(nums []int) int {
// 	return get(nums, 0, len(nums)-1).mSum
// }

// func pushUp(l, r Status) Status {
// 	iSum := l.iSum + r.iSum
// 	lSum := max(l.lSum, l.iSum+r.lSum)
// 	rSum := max(r.rSum, r.iSum+l.rSum)
// 	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
// 	return Status{lSum, rSum, mSum, iSum}
// }

// func get(nums []int, l, r int) Status {
// 	if l == r {
// 		return Status{nums[l], nums[l], nums[l], nums[l]}
// 	}
// 	m := (l + r) >> 1
// 	lSub := get(nums, l, m)
// 	rSub := get(nums, m+1, r)
// 	return pushUp(lSub, rSub)
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// type Status struct {
// 	lSum, rSum, mSum, iSum int
// }

func main() {
	test1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret1 := maxSubArray(test1)

	fmt.Println(ret1)

	test2 := []int{1}
	ret2 := maxSubArray(test2)

	fmt.Println(ret2)

	test3 := []int{5, 4, -1, 7, 8}
	ret3 := maxSubArray(test3)

	fmt.Println(ret3)

	fmt.Println(maxSubArray([]int{-4, -3, -1, -2}))
}
