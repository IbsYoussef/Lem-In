package main

import (
	"antfarm/lem-in/methods"

	"github.com/jroimartin/gocui"
)

func drawMap(x int, y int, lemin methods.LemIn, v3 *gocui.View) ([][]rune, [][]Point) {
	var direction [][]Point
	x = x - 1
	a := make([][]rune, y)
	for i := range a {
		a[i] = make([]rune, x)
	}
	i := 0
	for i < y {
		j := 0
		for j < x {
			a[i][j] = ' '
			j++
		}
		i++
	}
	i = 0
	for i < len(lemin.Start.Name) {
		if i == 0 {
			a[lemin.Start.Y*4][lemin.Start.X*6+i] = '['
		}
		a[lemin.Start.Y*4][lemin.Start.X*6+i+1] = rune(lemin.Start.Name[i])
		i++
		if i == len(lemin.Start.Name) {
			a[lemin.Start.Y*4][lemin.Start.X*6+i+1] = ']'
		}
	}
	i = 0
	for i < len(lemin.End.Name) {
		if i == 0 {
			a[lemin.End.Y*4][lemin.End.X*6+i] = '['
		}
		a[lemin.End.Y*4][lemin.End.X*6+i+1] = rune(lemin.End.Name[i])
		i++
		if i == len(lemin.End.Name) {
			a[lemin.End.Y*4][lemin.End.X*6+i+1] = ']'
		}
	}
	for _, room := range lemin.Rooms {
		i := 0
		for i < len(room.Name) {
			if i == 0 {
				a[room.Y*4][room.X*6+i] = '['
			}
			a[room.Y*4][room.X*6+i+1] = rune(room.Name[i])
			i++
			if i == len(room.Name) {
				a[room.Y*4][room.X*6+i+1] = ']'
			}
		}
	}
	_, possiblePath, result := methods.Calculate(&lemin)
	res := printSolution(lemin, possiblePath, result)
	for _, r := range res {
		i := 0
		tempDirection := []Point{}
		for i < len(result[r].Path)-1 {
			obstacles := findAllObsticles(a)
			startRoom := getRoom(result[r].Path[i].Name, lemin)
			start := Point{startRoom.X*6 + len(startRoom.Name) + 2, startRoom.Y * 4}
			endRoom := getRoom(result[r].Path[i+1].Name, lemin)
			end := Point{endRoom.X*6 - 1, endRoom.Y * 4}
			path := FindPath(start, end, obstacles)
			if path == nil {
				return nil, nil
			}
			for index, p := range path {
				if endRoom.Name == lemin.End.Name {
					if index != 0 && index != 1 && index != len(path)-1 && index != len(path)-2 {
						a[p.y][p.x] = '.'
					}
				} else {
					if index != 0 && index != len(path)-1 {
						a[p.y][p.x] = '.'
					}
				}
			}
			tempDirection = append(tempDirection, path...)
			i++
		}
		direction = append(direction, tempDirection)
	}
	return a, direction
}
func getRoom(name string, lemin methods.LemIn) methods.Room {
	for _, room := range lemin.Rooms {
		if room.Name == name {
			return room
		}
	}
	if name == lemin.Start.Name {
		return lemin.Start
	}
	if name == lemin.End.Name {
		return lemin.End
	}
	return methods.Room{}
}
func findAllObsticles(a [][]rune) []Point {
	var obs []Point
	for y, lines := range a {
		for x, line := range lines {
			if line != ' ' {
				obs = append(obs, Point{x: x, y: y})
			}
		}
	}
	return obs
}

func printSolution(lemin methods.LemIn, possiblePath [][]int, result []methods.LemIn) []int {
	var solutions []int
	i := 0
	for i < len(possiblePath) {
		ants := lemin.Ants
		turn := 0
		for methods.Move(possiblePath[i], ants, result, false) { // when last ant is not in end room
			if ants > 0 {
				methods.MoveFromStart(possiblePath[i], &ants, result, false)
			}
			turn++
		}
		solutions = append(solutions, turn)
		i++
	}
	min := solutions[0]
	indx := 0
	for index, m := range solutions {
		if min > m {
			min = m
			indx = index
		}
	}
	return possiblePath[indx] // fuster way
}
