package main

import (
	"container/list"
	"fmt"
)

func orangesRotting(grid [][]int) int {
	//bfs
	m := len(grid)
	n := len(grid[0])

	queue := list.New()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue.PushBack(Postion{i, j})
			}
		}
	}
	if queue.Len() > 0 {
		queue.PushBack(Postion{-1, -1})
	}

	minutes := 0
	for {
		if queue.Len() == 0 {
			break
		}
		e := queue.Front()
		p := e.Value.(Postion)
		queue.Remove(e)

		if p.x == -1 && p.y == -1 {
			minutes++
			if queue.Len() > 0 {
				queue.PushBack(Postion{-1, -1})
			}
			continue
		}

		if p.x-1 >= 0 {
			orangeRot(grid, queue, p.x-1, p.y)
		}

		if p.x+1 < m {
			orangeRot(grid, queue, p.x+1, p.y)
		}

		if p.y-1 >= 0 {
			orangeRot(grid, queue, p.x, p.y-1)
		}

		if p.y+1 < n {
			orangeRot(grid, queue, p.x, p.y+1)
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	if minutes > 0 {
		minutes--
	}

	return minutes

}

func orangeRot(grid [][]int, queue *list.List, x, y int) {
	if grid[x][y] == 1 {
		grid[x][y] = 2
		queue.PushBack(Postion{x, y})
	}
}

type Postion struct {
	x, y int
}

func main() {
	m1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(m1))

	m2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Println(orangesRotting(m2))

	m3 := [][]int{{0, 2}}
	fmt.Println(orangesRotting(m3))
}
