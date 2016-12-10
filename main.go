package main

import (
	"fmt"
	"io/ioutil"
)

var letters []string
var blue = "\033[0;34m"
var red = "\033[0;31m"
var AllCases = []Lcase{
	{
		Truth: false,
		Stype: "Redtext",
		Handler: func(red string) {
			fmt.Print(red)
		},
	},
	{
		Truth: true,
		Stype: "Bluetext",
		Handler: func(blue string) {
			fmt.Print(blue)
		},
	},
}

type Lcase struct {
	Truth   bool
	Stype   string
	Handler func(str string)
}

type Letter struct {
	Char  string
	Byte  byte
	Cases []Lcase
}

func main() {
	sortletters()
	f, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic("YOUR ALL LAME")
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
			if CurrentByte.Char == letters[q] {
				CurrentByte.Cases[q].Truth = true
			}
		}
		for w := 0; w < len(AllCases); w++ {
			if CurrentByte.Cases[w].Truth == true {
				CurrentByte.Cases[w].Handler(red)
				fmt.Print(CurrentByte.Char)
				fmt.Print("\033[0m")
			}
		}
	}
}

func sortletters() {
	for i := 33; i < 128; i++ {
		char := fmt.Sprintf("%c", i)
		letters = append(letters, string(char))
	}
}
