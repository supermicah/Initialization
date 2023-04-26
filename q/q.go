package main

import "fmt"

//m*n

/*
0 0 0 0
0 1 1 1
0 1 0 0
*/
type S struct {
	x int
	y int
}

func main() {
	n := make(map[int]map[int]int)
	n[0] = make(map[int]int)
	n[1] = make(map[int]int)
	n[2] = make(map[int]int)
	n[0][0] = 0
	n[0][1] = 0
	n[0][2] = 0
	n[0][3] = 0
	n[1][0] = 0
	n[1][1] = 1
	n[1][2] = 1
	n[1][3] = 1
	n[2][0] = 0
	n[2][1] = 1
	n[2][2] = 0
	n[2][3] = 0

	queue := make([]S, 0)
	queue = append(queue, S{x: 0, y: 0})
	count := 1
	visited := make(map[int]map[int]bool)
	xl := len(n)
	yl := len(n[0])
	for i := 0; i < xl; i++ {
		visited[i] = make(map[int]bool)
	}
	visited[0][0] = true

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		x := current.x
		y := current.y

		for i := 0; i < 4; i++ {
			sx := x + dx[i]
			sy := y + dy[i]
			if sx < xl && sx >= 0 && sy < yl && sy >= 0 {
				if _, ok := visited[sx][sy]; ok {
					continue
				}
				visited[sx][sy] = true
				if v, ok := n[sx][sy]; ok && v == 0 {
					println(sx, sy)
					queue = append(queue, S{x: sx, y: sy})
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
