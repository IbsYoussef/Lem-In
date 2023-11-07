package main

import (
	"antfarm/lem-in/methods"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"v1", "v3"}
	active  = 0
)

const NumGoroutines = 10

var (
	done = make(chan struct{})
	wg   sync.WaitGroup
	mu   sync.Mutex // protects ctr
	ctr  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}
func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]
	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}
	if nextIndex == 0 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}
	active = nextIndex
	return nil
}
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, maxX/4-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Task"
		v.Editable = true
		v.Wrap = true
		// var t string
		// scanner := bufio.NewScanner(os.Stdin)
		// for scanner.Scan() {
		// 	text := scanner.Text()
		// 	t += text + "\n"
		// }
		// fmt.Fprintln(v, t)
		//	_, possiblePath, result := calculate(&lemin)
		//	printing(lemin, possiblePath, result)
		if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("v2", maxX/4-1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Playground"
		v.Wrap = true
	}
	if v, err := g.SetView("v3", 0, maxY/2, maxX/4-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Log"
		v.Wrap = true
		v.Editable = true
	}
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func main() {
	lemin := methods.LemIn{}
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManagerFunc(layout)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
	var t string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		t += text + "\n"
	}
	methods.ParseFile(&lemin, t)
	for i := 0; i < NumGoroutines; i++ {
		wg.Add(1)
		go counter(g, lemin, t)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
func counter(g *gocui.Gui, lemin methods.LemIn, t string) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case <-time.After(500 * time.Millisecond):
			mu.Lock()
			n := ctr
			ctr++
			mu.Unlock()
			if n < 30 {
				g.Update(func(g *gocui.Gui) error {
					v, _ := g.View("v1")
					v.Clear()
					fmt.Fprintln(v, t)
					v, err := g.View("v2")
					if err != nil {
						return err
					}
					v.Clear()
					x, y := v.Size()
					v3, _ := g.View("v3")
					a, path := drawMap(x, y, lemin, v3)
					if a == nil {
						return nil
					}
					printMapJustRooms(v, a, path)
					return nil
				})
			} else if n > 30 && n < 60 {
				g.Update(func(g *gocui.Gui) error {
					v, _ := g.View("v1")
					v.Clear()
					fmt.Fprintln(v, t)
					v, err := g.View("v2")
					if err != nil {
						return err
					}
					v.Clear()
					x, y := v.Size()
					v3, _ := g.View("v3")
					a, path := drawMap(x, y, lemin, v3)
					if a == nil {
						return nil
					}
					printMapJustRoomsAndDots(v, a, path)
					return nil
				})
			} else if n > 60 && n < 150 {
				g.Update(func(g *gocui.Gui) error {
					v, _ := g.View("v1")
					v.Clear()
					fmt.Fprintln(v, t)
					v, err := g.View("v2")
					if err != nil {
						return err
					}
					v.Clear()
					x, y := v.Size()
					v3, _ := g.View("v3")
					a, path := drawMap(x, y, lemin, v3)
					if a == nil {
						return nil
					}
					printMap(v, a, path)
					return nil
				})
			}
		}
	}
}
func printMap(v *gocui.View, a [][]rune, path [][]Point) {
	for y, line := range a {
		l := ""
		for x, c := range line {
			if c == '.' {
				n := whichPath(x, y, path)
				l += "\033[3" + strconv.Itoa(n+1) + ";1m" + string(c) + "\033[0m"
			} else {
				l += string(c)
			}
		}
		fmt.Fprintln(v, l)
	}
}
func printMapJustRooms(v *gocui.View, a [][]rune, path [][]Point) {
	for _, line := range a {
		l := ""
		for _, c := range line {
			if c == '.' {
				l += " "
			} else {
				l += string(c)
			}
		}
		fmt.Fprintln(v, l)
	}
}
func printMapJustRoomsAndDots(v *gocui.View, a [][]rune, path [][]Point) {
	for _, line := range a {
		l := ""
		for _, c := range line {
			l += string(c)
		}
		fmt.Fprintln(v, l)
	}
}
func whichPath(x int, y int, path [][]Point) int {
	var n int
	for index, p := range path {
		n = index
		for _, point := range p {
			if point.x == x && point.y == y {
				return n
			}
		}
	}
	return -1
}
