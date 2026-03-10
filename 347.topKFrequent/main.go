package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// 堆排序
	// if k > len(nums) {
	// 	k = len(nums)
	// }
	// num_map := make(map[int]int, 0)
	// for _, v := range nums {
	// 	num_map[v]++
	// }

	// //fmt.Println(num_map)

	// heap := make([]int, k)
	// index := 0
	// for key, v := range num_map {
	// 	if index < k {
	// 		heap[index] = key
	// 	} else if index >= k {
	// 		if index == k {
	// 			for i := k/2 - 1; i >= 0; i-- {
	// 				shifDown(heap, i, num_map)
	// 			}
	// 		}
	// 		if v > num_map[heap[0]] {
	// 			heap[0] = key
	// 			shifDown(heap, 0, num_map)
	// 		}
	// 	}
	// 	index++
	// }

	// return heap

	// 桶排序
	// num_map := make(map[int]int, 0)
	// for _, v := range nums {
	// 	num_map[v]++
	// }

	// l := len(nums)
	// //fmt.Println(num_map)
	// bucket := make([][]int, l+1)
	// for key, v := range num_map {
	// 	bucket[v] = append(bucket[v], key)
	// }

	// ret := make([]int, 0, k)
	// for i := len(bucket) - 1; i >= 0; i-- {
	// 	if bucket[i] != nil {
	// 		ret = append(ret, bucket[i]...)
	// 	}

	// 	if len(ret) >= k {
	// 		break
	// 	}
	// }
	// return ret[:k]

	//快速选择
	num_map := make(map[int]int, 0)
	for _, v := range nums {
		num_map[v]++
	}

	l := len(num_map)

	Freqs := make([]*Freq, 0, l)
	for key, v := range num_map {
		Freqs = append(Freqs, &Freq{key, v})
	}

	quick_select(Freqs, 0, l-1, l-k)

	ret := make([]int, 0, k)
	for i := l - 1; i >= 0 && len(ret) < k; i-- {
		ret = append(ret, Freqs[i].value)
	}
	return ret

}

func quick_select(freqs []*Freq, l, r, k int) {
	if l == r {
		return
	}

	point_cnt := freqs[l].cnt
	i := l - 1
	j := r + 1

	for i < j {
		for i++; freqs[i].cnt < point_cnt; i++ {

		}

		for j--; freqs[j].cnt > point_cnt; j-- {

		}

		if i < j {
			freqs[i], freqs[j] = freqs[j], freqs[i]
		}
	}

	if j < k {
		quick_select(freqs, j+1, r, k)
	} else {
		quick_select(freqs, l, j, k)
	}

}

type Freq struct {
	value, cnt int
}

func shifDown(head []int, i int, num_map map[int]int) {
	l := len(head)
	for {
		left_child := 2*i + 1
		right_child := left_child + 1
		if left_child < l {
			swap_postion := left_child
			if right_child < l {
				if num_map[head[left_child]] > num_map[head[right_child]] {
					swap_postion = right_child
				}
			}

			if num_map[head[i]] > num_map[head[swap_postion]] {
				head[i], head[swap_postion] = head[swap_postion], head[i]
				i = swap_postion
			} else {
				break
			}
		} else {
			break
		}
	}

}
func main() {
	test1 := []int{1, 1, 1, 2, 2, 3}
	ret1 := topKFrequent(test1, 2)

	fmt.Println(ret1)

	test2 := []int{1}
	ret2 := topKFrequent(test2, 1)

	fmt.Println(ret2)

	test3 := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	ret3 := topKFrequent(test3, 2)

	fmt.Println(ret3)

	fmt.Println("ret4", topKFrequent([]int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5}, 7))
}
