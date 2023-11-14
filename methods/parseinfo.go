package methods

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func ParseFile(lemin *LemIn, s string) error {
	var err error
	var temp []string
	var k int
	isFirstTime := true
	str := strings.Split(s, "\n")
	l := len(str)
	if l > 1 && str[l-1] == "" { // if las element is empty line
		k = 1
	} else {
		k = 0
	}
	i := 0
	for i < l-k {
		if str[i] == "##start" { // add start room
			i++
			err = parseRoom(lemin, str[i], 1)
			if err != nil {
				return errors.New("ERROR: invalid data format, invalid start room")
			}
		} else if str[i] == "##end" { // add end room
			i++
			err = parseRoom(lemin, str[i], -1)
			if err != nil {
				return errors.New("ERROR: invalid data format, invalid end room")
			}
		} else if len(str[i]) == 0 {
			return errors.New("ERROR: invalid data format, file contain empty string")
		} else if str[i][0] != '#' { // if a comment skip
			if isFirstTime {
				lemin.Ants, err = strconv.Atoi(str[i])
				if err != nil || lemin.Ants <= 0 {
					return errors.New("ERROR: invalid data format, invalid number of Ants")
				}
				isFirstTime = false
			} else {
				temp = strings.Split(str[i], " ")
				if len(temp) == 3 {
					err = parseRoom(lemin, str[i], 0) // add a room
					if err != nil {
						return err
					}
				} else if len(temp) == 1 {
					err = parseLink(lemin, str[i]) // add a link
					if err != nil {
						return err
					}
				} else {
					return errors.New("ERROR: invalid data format")
				}
			}
		}
		i++
	}
	err = isAllCredentialsValid(lemin)
	if err != nil {
		return err
	}
	return nil
}

func isAllCredentialsValid(lemin *LemIn) error {
	if len(lemin.Rooms) == 0 {
		return errors.New("ERROR: invalid data format, no rooms")
	}
	if len(lemin.Links) == 0 {
		return errors.New("ERROR: invalid data format, no links")
	}
	if lemin.Start.Name == "" || lemin.End.Name == "" {
		return errors.New("ERROR: invalid data format, no start or end room")
	}
	return nil
}

func checkDublicatesInRooms(lemin *LemIn) error {
	count := 0
	for _, name := range lemin.Rooms {
		count = 0
		for _, name2 := range lemin.Rooms {
			if name.Name == name2.Name || name.Name == lemin.Start.Name || name.Name == lemin.End.Name {
				count++
			}
			if count == 2 {
				return errors.New("ERROR: invalid data format, dublicates in rooms")
			}
		}
	}
	return nil
}
func checkLinks(lemin *LemIn) error {
	l := lemin.Links
	for _, link := range l {
		if !isRoomExist(*lemin, link.Room1) || !isRoomExist(*lemin, link.Room2) { //function that check is every room exist in links
			return errors.New("ERROR: can't create tunnel")
		}
		removeDublicates(lemin, link)
	}
	return nil
}
func removeDublicates(lemin *LemIn, link Link) {
	count := 0
	var links []int
	for index, l := range lemin.Links {
		if (l.Room1.Name == link.Room1.Name && l.Room2.Name == link.Room2.Name) || (l.Room1.Name == link.Room2.Name && l.Room2.Name == link.Room1.Name) {
			count++
		}
		if count == 2 {
			if count == 2 {
				count--
			}
			links = append(links, index)
		}
	}
	sort.Ints(links)
	removeLink(lemin, links)
}
func removeLink(lemin *LemIn, indexs []int) {
	var temp []Link
	for i, link := range lemin.Links {
		if !contains(indexs, i) {
			temp = append(temp, link)
		}
	}
	(*lemin).Links = temp
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func isRoomExist(lemin LemIn, room Room) bool {
	for _, r := range lemin.Rooms {
		if r.Name == room.Name {
			return true
		}
	}
	if room.Name == lemin.Start.Name || room.Name == lemin.End.Name {
		return true
	}
	return false
}
func parseRoom(lemin *LemIn, s string, flag int) error {
	room := Room{}
	if len(s) == 0 {
		return errors.New("ERROR: invalid data format, file contain empty string")
	}
	str := strings.Split(s, " ")
	if len(str) != 3 {
		return errors.New("ERROR: invalid data format, invalid room")
	}
	a, err := strconv.Atoi(str[1])
	if err != nil {
		return errors.New("ERROR: invalid data format, invalid room")
	}
	b, err := strconv.Atoi(str[2])
	if err != nil {
		return errors.New("ERROR: invalid data format, invalid room")
	}
	room.Name = str[0]
	if len(room.Name) == 0 || room.Name[0] == 'L' || strings.Contains(room.Name, "-") {
		return errors.New("ERROR: invalid data format, invalid room name")
	}
	room.X = a
	room.Y = b
	if flag == 1 {
		lemin.Start = room
	} else if flag == -1 {
		lemin.End = room
	} else {
		lemin.Rooms = append(lemin.Rooms, room)
	}
	err = checkDublicatesInRooms(lemin)
	if err != nil {
		return err
	}
	return nil
}
func parseLink(lemin *LemIn, s string) error {
	link := Link{}
	str := strings.Split(s, "-")
	if len(str) != 2 {
		return errors.New("ERROR: invalid data format, invalid link")
	}
	link.Room1.Name = str[0]
	link.Room2.Name = str[1]
	lemin.Links = append(lemin.Links, link)
	err := checkLinks(lemin) //function that check is every room exist in links also dublicates
	if err != nil {
		return err
	}
	return nil
}
