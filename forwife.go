package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	messages := []string{
		"Good morning my favorite human ❤️",
	"Just a reminder: You are loved a lot",
	"I’d pick you in every universe",
	"Smile, you make the world 100x better",
		"I'm so lucky to call you mine",
	}

	heart := `
   ***     ***
  *****   *****
 ******* *******
  *************
   ***********
     *******
       ***
        *
`

	fmt.Println("Hey beautiful 💕")
	fmt.Println(messages[rand.Intn(len(messages))])
	fmt.Println(heart)
	fmt.Println("Made by your programmer husband who DOES have a life... it's you :)")
}
