package main

import "fmt"

type MedianFinder struct {
	minHeap    []int
	minHeapLen int
	maxHeap    []int
	maxHeapLen int
}

func Constructor() MedianFinder {
	return *new(MedianFinder)
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeapLen == this.minHeapLen {
		if this.maxHeapLen > 0 && this.maxHeap[0] < num {
			this.minHeap = append(this.minHeap, num)
			this.shiftUp(this.minHeap, this.minHeapLen, true)
			this.minHeapLen++
			return
		}
		this.maxHeap = append(this.maxHeap, num)
		this.shiftUp(this.maxHeap, this.maxHeapLen, false)
		this.maxHeapLen++
	} else {
		if this.maxHeapLen > this.minHeapLen {
			if this.maxHeapLen > 0 {
				if num < this.maxHeap[0] {
					num, this.maxHeap[0] = this.maxHeap[0], num
					//fmt.Println("shiftDown", this.maxHeap, this.maxHeapLen)
					this.shiftDown(this.maxHeap, 0, this.maxHeapLen-1, false)
				}
			}
			this.minHeap = append(this.minHeap, num)
			this.shiftUp(this.minHeap, this.minHeapLen, true)
			this.minHeapLen++
		} else {
			if this.minHeapLen > 0 {
				if num > this.minHeap[0] {
					num, this.minHeap[0] = this.minHeap[0], num
					this.shiftDown(this.minHeap, 0, this.minHeapLen-1, true)
				}
			}
			this.maxHeap = append(this.maxHeap, num)
			this.shiftUp(this.maxHeap, this.maxHeapLen, false)
			this.maxHeapLen++
		}

	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeapLen <= 0 {
		return 0
	}
	if this.maxHeapLen == this.minHeapLen {
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2
	} else if this.maxHeapLen > this.minHeapLen {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}

func (this *MedianFinder) shiftUp(heap []int, i int, isMinHeap bool) {
	for {
		parent := (i - 1) / 2
		if parent >= 0 {
			if (isMinHeap && heap[parent] > heap[i]) || (!isMinHeap && heap[parent] < heap[i]) {
				heap[parent], heap[i] = heap[i], heap[parent]
				i = parent
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (this *MedianFinder) shiftDown(heap []int, i, end int, isMinHeap bool) {
	for {
		left_child := 2*i + 1
		if left_child <= end {
			right_child := left_child + 1
			swap_postion := left_child
			if right_child <= end {
				if (isMinHeap && heap[right_child] < heap[left_child]) || (!isMinHeap && heap[right_child] > heap[left_child]) {
					swap_postion = right_child
				}
			}

			if (isMinHeap && heap[i] > heap[swap_postion]) || (!isMinHeap && heap[i] < heap[swap_postion]) {
				heap[i], heap[swap_postion] = heap[swap_postion], heap[i]
				i = swap_postion
			} else {
				break
			}
		} else {
			break
		}

	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	obj := Constructor()
	// obj.AddNum(-1)
	// fmt.Println(obj.maxHeap, obj.minHeap)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-2)
	// fmt.Println(obj.maxHeap, obj.minHeap)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-3)
	// fmt.Println(obj.maxHeap, obj.minHeap)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-4)
	// fmt.Println(obj.maxHeap, obj.minHeap)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-5)
	// fmt.Println(obj.maxHeap, obj.minHeap)
	// fmt.Println(obj.FindMedian())

	obj.AddNum(1)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(2)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(4)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(5)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(6)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(7)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(8)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(9)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(10)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())

	//fmt.Println(obj.heap)
}
