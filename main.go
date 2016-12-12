package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

var letters []string
var blue = "\033[0;34m"
var red = "\033[0;31m"
var AllCases = []Lcase{
	{
		Test: func(ltr byte) bool {
			if unicode.IsLower([]rune(string(ltr))[0]) == true {
				return true
			}
			return false
		},
		Truth: false,
		Stype: "Redtext",
		Handler: func() {
			fmt.Print(red)
		},
	},
	{
		Test: func(ltr byte) bool {
			if unicode.IsUpper([]rune(string(ltr))[0]) == true {
				return true
			}
			return false
		},
		Truth: false,
		Stype: "Bluetext",
		Handler: func() {
			fmt.Print(blue)
		},
	},
}

type Lcase struct {
	Test    func(ltr byte) bool
	Truth   bool
	Stype   string
	Handler func()
}

type Letter struct {
	Char  string
	Byte  byte
	Cases []Lcase
}

func main() {
	f, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic("YOU'RE ALL LAME")
	}
	//raw := string(bytes.Trim(f, "\x00"))
	pcases := AllCases
	for i := 0; i < len(f); i++ {
		CurrentByte := Letter{
			Char:  string(f[i]),
			Byte:  f[i],
			Cases: pcases,
		}
		//kek
		for q := 0; q < len(CurrentByte.Cases); q++ {
			truth := CurrentByte.Cases[q].Test(CurrentByte.Byte)
			CurrentByte.Cases[q].Truth = truth
		}
		caseintercept := false
		for w := 0; w < len(AllCases); w++ {
			if CurrentByte.Cases[w].Truth == true {
				caseintercept = true
				CurrentByte.Cases[w].Handler()
				fmt.Print(CurrentByte.Char)
				fmt.Print("\033[0m")
				break
			}
		}
		if caseintercept == false {
			fmt.Print("\033[0m")
			fmt.Print(CurrentByte.Char)
		}
	}
}
