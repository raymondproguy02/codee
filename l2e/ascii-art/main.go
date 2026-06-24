package main

import (
	"ascii-art/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	if len(os.Args[1]) == 0 {
		return
	}

	valid, err := parser.Val(os.Args[1])
	if err != nil {
		fmt.Printf("invalid %c", valid)
		return
	}
	font, err := parser.LoadBanner(os.Args[2])
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	words := parser.Split(os.Args[1])
	res := parser.GenArt(font, words)
	fmt.Println(res)
}
