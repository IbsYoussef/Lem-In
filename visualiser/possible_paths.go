package main

import (
	"container/heap"
	"math"
)

// Тип Point представляє точку на двовимірній сітці.
type Point struct {
	x, y int
}

// Тип Node представляє вузол графа для алгоритму A*.
type Node struct {
	point   Point
	f, g, h float64
	parent  *Node
}

// Тип PriorityQueue представляє пріоритетну чергу для вузлів.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Функція ManhattanDistance обчислює відстань Манхеттен між двома точками.
func ManhattanDistance(p1, p2 Point) float64 {
	return math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y))
}

// Функція FindPath знаходить найкоротший шлях між двома точками з врахуванням можливих перешкод.
func FindPath(start, end Point, obstacles []Point) []Point {
	openSet := make(PriorityQueue, 0)
	closedSet := make(map[Point]bool)
	startNode := &Node{
		point:  start,
		f:      0,
		g:      0,
		h:      ManhattanDistance(start, end),
		parent: nil,
	}
	heap.Push(&openSet, startNode)
	for len(openSet) > 0 {
		currentNode := heap.Pop(&openSet).(*Node)
		if currentNode.point == end {
			return reconstructPath(currentNode)
		}
		closedSet[currentNode.point] = true
		neighbors := getNeighbors(currentNode.point)
		for _, neighbor := range neighbors {
			if isObstacle(neighbor, obstacles) || closedSet[neighbor] {
				continue
			}
			gScore := currentNode.g + 1
			found := false
			neighborNode := &Node{point: neighbor}
			for _, node := range openSet {
				if node.point == neighbor {
					neighborNode = node
					found = true
					break
				}
			}
			if !found || gScore < neighborNode.g {
				neighborNode.g = gScore
				neighborNode.h = ManhattanDistance(neighbor, end)
				neighborNode.f = neighborNode.g + neighborNode.h
				neighborNode.parent = currentNode
				if !found {
					heap.Push(&openSet, neighborNode)
				}
			}
		}
	}
	// Якщо неможливо знайти шлях
	return nil
}
func getNeighbors(point Point) []Point {
	// Функція, яка повертає сусідні точки для даної точки (зазвичай 4 або 8 сусідів).
	// Ваша реалізація може залежати від конкретної конфігурації сітки.
	// В цьому прикладі реалізована конфігурація з 4 сусідами.
	neighbors := []Point{
		{point.x - 1, point.y},
		{point.x + 1, point.y},
		{point.x, point.y - 1},
		{point.x, point.y + 1},
	}
	return neighbors
}
func isObstacle(point Point, obstacles []Point) bool {
	// Функція перевіряє, чи точка є перешкодою.
	for _, obstacle := range obstacles {
		if point == obstacle {
			return true
		}
	}
	return false
}
func reconstructPath(node *Node) []Point {
	// Функція відновлює шлях з кінцевого вузла до початкового, проходячи по батькам вузлів.
	path := []Point{node.point}
	for node.parent != nil {
		node = node.parent
		path = append([]Point{node.point}, path...)
	}
	return path
}
