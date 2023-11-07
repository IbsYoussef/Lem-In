package main

import (
	"antfarm/lem-in/methods"
	"fmt"
	"os"
)

func main() {
	lemin := methods.LemIn{}

	args := os.Args
	if len(args) == 2 {
		s, err := methods.OpenFile(args[1])
		if err != nil {
			fmt.Println("ERROR: can't open:", args[1])
			return
		}
		err = methods.ParseFile(&lemin, s)
		if err != nil {
			fmt.Printf("%v", err.Error())
			return
		}

		err, possibelPaths, result := methods.Calculate(&lemin)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		methods.PrintGrahpMoves(lemin, possibelPaths, result)
	} else {
		fmt.Println("Error: Invalid data format, File Example Usage: example01.txt")
	}

}
