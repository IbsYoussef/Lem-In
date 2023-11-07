package methods

import (
	"errors"
	"sort"
)

func Calculate(lemin *LemIn) (error, [][]int, []LemIn) {
	// go until no more tunnels forward or is end room
	// when we go we remove rooms and conection
	// start from start room
	//then go to connected rooms, destroy conections of previous start and destroy previous start room
	// set new start room
	// if last elem in path is end room or if can go forward from last room in path go to next elem
	// || isPathLastElemIsEndRoom(lemins) when exit to next elem
	var lemins []LemIn
	var result []LemIn
	lemins = append(lemins, *lemin)
	for len(lemins) != 0 {
		temp := lemins[0]
		removeFirst(&lemins) // remove 0 element in lemins
		if temp.Start.Name != lemin.End.Name {
			nextLemins := findNextRooms(temp) // inside need to remove connection and rooms
			if len(nextLemins) != 0 {         // go to next elem if nextlemint == 0
				lemins = append(lemins, nextLemins...) // add nextLemins to lemins
			}
		} else {
			temp.Path = append(temp.Path, lemin.End)
		}
		result = append(result, temp)
	}
	clearingResultsPath(&result, lemin.End)
	if len(result) == 0 {
		return errors.New("ERROR: start room not connected to end room"), nil, nil
	}
	uniqPath := createUniquePaths(result)
	possiblePath := possiblePath(uniqPath, result)
	return nil, possiblePath, result
}

func createUniquePaths(result []LemIn) [][]int {
	var uniq [][]int
	var tempComb []int
	var comb [][]int
	for index := range result {
		if len(result[0].Rooms) > 20 || len(result) > 13 {
			if index < 10 {
				uniq = append(uniq, []int{index})
			}
		} else {
			uniq = append(uniq, []int{index})
		}
	}
	for index := range result {
		if len(result[0].Rooms) > 20 || len(result) > 13 {
			if index < 10 {
				tempComb = append(tempComb, index)
			}
		} else {
			tempComb = append(tempComb, index)
		}
	}
	comb = uniq
	i := 1
	for i < len(tempComb) {
		comb = findCombination(tempComb, comb) // after need to add to unique
		uniq = append(uniq, comb...)
		i++
	}
	return uniq
}

func findCombination(tempComb []int, temptempComb [][]int) [][]int {
	var comb [][]int
	start := 0
	l := len(tempComb)
	for start < len(temptempComb) {
		i := 0
		for i < l {
			temp1 := coppy(temptempComb[start])
			if !isContains(temp1, tempComb[i], comb) {
				//temp := temptempComb[start]
				temp := coppy(temptempComb[start])
				//	fmt.Println(tempComb[i])
				//	fmt.Println(temp)
				temp = append(temp, tempComb[i])
				//	fmt.Println(temp)
				comb = append(comb, temp)
			}
			i++
		}
		start++
	}
	return comb
}

func coppy(temp []int) []int {
	var result []int
	result = append(result, temp...)
	return result
}

func isContains(list []int, k int, uniq [][]int) bool {
	end := 0
	l := len(list)
	for end < len(uniq) {
		//	temp := list
		temp := append(list, k)
		if cont(uniq[end], temp) {
			return true
		}
		end++
	}
	i := 0
	for i < l {
		if list[i] == k {
			return true
		}
		i++
	}
	return false
}

func cont(list []int, c []int) bool {
	sort.Ints(list)
	sort.Ints(c)
	i := 0
	for i < len(list) {
		if list[i] != c[i] {
			return false
		}
		i++
	}
	return true
}

func clearingResultsPath(result *[]LemIn, end Room) {
	var newResult []LemIn
	for _, lem := range *result {
		if lem.Path != nil && lem.Path[len(lem.Path)-1].Name == end.Name {
			newResult = append(newResult, lem)
		}
	}
	*result = newResult
}

func findNextRooms(temp LemIn) []LemIn {
	var result []LemIn
	var tempName string
	for index, link := range temp.Links {
		if temp.Start.Name == link.Room1.Name || temp.Start.Name == link.Room2.Name {
			if temp.Start.Name == link.Room1.Name {
				tempName = link.Room2.Name
			} else {
				tempName = link.Room1.Name
			}
			if !isRoomAlreadyinPath(temp.Path, temp.Start) && !isRoomAlreadyinPath(temp.Path, temp.End) {
				t := temp
				t.Path = append(t.Path, t.Start)
				t.Start.Name = tempName
				removeLink(&t, []int{index})
				result = append(result, t)
			}
		}
	}
	return result
}

func isRoomAlreadyinPath(path []Room, start Room) bool {
	for _, room := range path {
		if room.Name == start.Name {
			return true
		}
	}
	return false
}

func removeFirst(lemins *[]LemIn) {
	var temp []LemIn
	for index, lem := range *lemins {
		if index != 0 {
			temp = append(temp, lem)
		}
	}
	*lemins = temp
}

// 1 2 3 4
// {[1,2], [1,3], [1,4],[2,3], [2,4], [3,4]}
// {[1,2,3],[1,2,4], [2,3,4], [1,3,4]}
// {[1,2,3,4]}
