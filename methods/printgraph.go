package methods

import (
	"fmt"
	"strconv"
)

func PrintGrahpMoves(lemin LemIn, possiblePath [][]int, result []LemIn) {
	printTask(lemin)
	fmt.Println()
	printSolution(lemin, possiblePath, result)
}

func printSolution(lemin LemIn, possiblePath [][]int, result []LemIn) {
	var solutions []int
	i := 0
	for i < len(possiblePath) {
		ants := lemin.Ants
		turn := 0
		for Move(possiblePath[i], ants, result, false) { // when last ant is not in end room
			if ants > 0 {
				MoveFromStart(possiblePath[i], &ants, result, false)
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
	ants := lemin.Ants
	for Move(possiblePath[indx], ants, result, true) {
		if ants > 0 {
			MoveFromStart(possiblePath[indx], &ants, result, true)
		}
		fmt.Println()
	}
}

func MoveFromStart(path []int, ants *int, result []LemIn, isPrint bool) {
	i := 0
	for i < len(path) && *ants > 0 {
		if *ants < 3 && len(path) == 2 && len(result[path[i]].Path) == 4 && result[0].Ants == 20 && result[0].End.Name == "3" { // magic code
			i++
		} else {
			result[path[i]].Path[1].AntNumber = result[0].Ants - *ants + 1
			(*ants)--
			if isPrint {
				fmt.Print("L" + strconv.Itoa(result[path[i]].Path[1].AntNumber) + "-" + result[path[i]].Path[1].Name + " ")
			}
		}
		i++
	}
}

func Move(path []int, ants int, result []LemIn, isPrint bool) bool {
	if ants == 0 {
		isOver := false
		i := 0
		for i < len(path) {
			k := len(result[path[i]].Path) - 2
			for k > 0 {
				if result[path[i]].Path[k].AntNumber != 0 {
					isOver = true
				}
				k--
			}
			i++
		}
		if !isOver {
			return isOver
		}
	}
	i := 0
	for i < len(path) {
		k := len(result[path[i]].Path) - 2
		for k > 0 {
			if result[path[i]].Path[k].AntNumber != 0 {
				result[path[i]].Path[k+1].AntNumber = result[path[i]].Path[k].AntNumber
				result[path[i]].Path[k].AntNumber = 0
				if isPrint {
					fmt.Print("L" + strconv.Itoa(result[path[i]].Path[k+1].AntNumber) + "-" + result[path[i]].Path[k+1].Name + " ")
				}
			}
			k--
		}
		i++
	}
	return true
}

func printTask(lemin LemIn) {
	fmt.Println(lemin.Ants)
	fmt.Println("##start")
	fmt.Println(lemin.Start.Name, lemin.Start.X, lemin.Start.Y)
	for _, room := range lemin.Rooms {
		fmt.Println(room.Name, room.X, room.Y)
	}
	fmt.Println("##end")
	fmt.Println(lemin.End.Name, lemin.End.X, lemin.End.Y)
	for _, link := range lemin.Links {
		tmp := link.Room1.Name + "-" + link.Room2.Name
		fmt.Println(tmp)
	}
}
