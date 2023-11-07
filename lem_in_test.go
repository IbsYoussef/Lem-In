package main

import (
	"antfarm/lem-in/methods"
	"testing"
)

// str means string before and the expected stands for the 'result we expect'
type LemInTest struct {
	file string
}

var parsingTests = []LemInTest{
	{"examples/example00.txt"},
	{"examples/example01.txt"},
	{"examples/example02.txt"},
	{"examples/example03.txt"},
	{"examples/example04.txt"},
	{"examples/example05.txt"},
	{"examples/example06.txt"},
	{"examples/example07.txt"},
	{"examples/example08.txt"},
	{"examples/example09.txt"},
	{"examples/example11.txt"},
	{"examples/example12.txt"},
	{"examples/example13.txt"},
	{"examples/example14.txt"},
	{"examples/example15.txt"},
	{"examples/example16.txt"},
}

func TestParsing(t *testing.T) {
	for _, test := range parsingTests {
		lemin := methods.LemIn{}
		s, err := methods.OpenFile(test.file)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		err = methods.ParseFile(&lemin, s)
		if err != nil {
			t.Errorf(test.file + "\n")
			t.Errorf(err.Error())
			return
		}
	}
}
