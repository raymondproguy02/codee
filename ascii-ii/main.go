package main

import (
	"code/perser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		if len(os.Args[1]) == 0 {
			return
		}
		font, err := perser.Load(os.Args[2])
		if err != nil {
			fmt.Println("err", err)
			return
		}
		lines := perser.Split(os.Args[1])
		res := perser.Generate(font, lines)
		fmt.Println(res)
	}
}
