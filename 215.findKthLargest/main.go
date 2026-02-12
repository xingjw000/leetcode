package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	// l := len(nums)
	// if k > l {
	// 	panic("k is larger than array len!")
	// }
	// k_nums := make([]int, 0, k)
	// for i := 0; i < k; i++ {
	// 	k_nums = append(k_nums, nums[i])
	// 	j := len(k_nums) - 1
	// 	for j > 0 {
	// 		if k_nums[j] < k_nums[(j-1)/2] {
	// 			swap(&k_nums[j], &k_nums[(j-1)/2])
	// 		} else {
	// 			break
	// 		}

	// 		j = (j - 1) / 2
	// 	}
	// }

	// // fmt.Println(k_nums)

	// for i := k; i < l; i++ {
	// 	j := 0
	// 	if nums[i] > k_nums[j] {
	// 		k_nums[j] = nums[i]
	// 		to_swap_index := 2*j + 1
	// 		for to_swap_index < k {
	// 			to_swap_nums := k_nums[to_swap_index]
	// 			if to_swap_index+1 < k {
	// 				if to_swap_nums > k_nums[to_swap_index+1] {
	// 					to_swap_index = to_swap_index + 1
	// 				}
	// 			}

	// 			if k_nums[j] > k_nums[to_swap_index] {
	// 				swap(&k_nums[j], &k_nums[to_swap_index])
	// 				j = to_swap_index
	// 				to_swap_index = 2*j + 1
	// 			} else {
	// 				break
	// 			}

	// 		}
	// 	}
	// 	//fmt.Println(k_nums)
	// }
	// //fmt.Println(k_nums)

	// return k_nums[0]

	// l := len(nums)
	// if k > l {
	// 	panic("k is larger than array len!")
	// }
	// k_nums := make([]int, k)
	// copy(k_nums, nums[:k])
	// buildMinHeap(k_nums)

	// for i := k; i < l; i++ {
	// 	if nums[i] > k_nums[0] {
	// 		k_nums[0] = nums[i]
	// 		siftDown(k_nums, 0, k)
	// 	}

	// }

	// return k_nums[0]

	target := len(nums) - k
	left := 0
	right := len(nums) - 1

	return partition(nums, left, right, target)

}

// func swap(a, b *int) {
// 	tmp := *a
// 	*a = *b
// 	*b = tmp
// }

// func buildMinHeap(heap []int) {
// 	l := len(heap)
// 	for i := l/2 - 1; i >= 0; i-- {
// 		siftDown(heap, i, l)
// 	}
// }

// func siftDown(heap []int, i int, size int) {
// 	for {
// 		left := 2*i + 1
// 		if left >= size {
// 			break
// 		}
// 		right := left + 1
// 		smallest := left
// 		if right < size && heap[smallest] > heap[right] {
// 			smallest = right
// 		}

// 		if heap[i] <= heap[smallest] {
// 			break
// 		}

// 		heap[i], heap[smallest] = heap[smallest], heap[i]
// 		i = smallest
// 	}
// }

func partition(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	randidx := rand.Intn(right-left+1) + left
	nums[randidx], nums[left] = nums[left], nums[randidx]
	pivot := nums[left]
	i := left - 1
	j := right + 1

	for i < j {
		for i++; nums[i] < pivot; {
			i++
		}

		for j--; nums[j] > pivot; {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		return partition(nums, left, j, k)
	} else {
		return partition(nums, j+1, right, k)
	}
}

func main() {
	test1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	ret1 := findKthLargest(test1, 4)

	fmt.Println(ret1)

	test2 := []int{3, 2, 1, 5, 6, 4}
	ret2 := findKthLargest(test2, 2)

	fmt.Println(ret2)

	// test3 := []int{-2, 1, 0, -2}
	// ret3 := subarraySum(test3, -1)

	// fmt.Println(ret3)
}
