package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // LoadBanner the function that reads from the banner file
// func LoadBanner(file string) ([]string, error) {
// 	// Remove extra extentions if provided by the user
// 	file = strings.TrimSuffix(file, ".txt")
// 	data, err := os.ReadFile("banners/" + file + ".txt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(data) == 0 {
// 		return nil, fmt.Errorf("Empty banner file")
// 	}
// 	var lines = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
// 	if len(lines) < 855 || len(lines) > 856 {
// 		return nil, fmt.Errorf("Incomplete banner file")
// 	}
// 	return lines, nil
// }
// package perser

// import "strings"

// func GenerateArt(fonts []string, words []string) string {
// 	var res strings.Builder
// 	for index, word := range words {
// 		if word == "" {
// 			res.WriteString("\n")
// 			continue
// 		}
// 		for i := 1; i <= 8; i++ {
// 			for _, ch := range word {
// 				res.WriteString(fonts[i+(int(ch-32)*9)])
// 			}
// 			res.WriteString("\n")
// 		}
// 		if index < len(words)-1 && words[index+1] != "" {
// 			res.WriteString("\n")
// 		}
// 	}
// 	return res.String()
// }
// package perser

// import (
// 	"fmt"
// 	"strings"
// )

// func SplitInput(input string) []string {
// 	input = strings.ReplaceAll(input, "\r\n", "\n")
// 	return strings.Split(input, "\\n")
// }

// func Validate(input string) (rune, error) {
// 	for _, r := range input {
// 		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
// 			return r, fmt.Errorf("Invalid character")
// 		}
// 	}
// 	return 0, nil
// }
// func main() {
// 	if len(os.Args) == 3 {
// 		if len(os.Args[1]) == 0 {
// 			return
// 		}
// 		fontFileLines, err := LoadBanner(os.Args[2])
// 		if err != nil {
// 			fmt.Printf("Error: %v", err)
// 			return
// 		}
// 		words := SplitInput(os.Args[1])
// 		result := GenerateArt(fontFileLines, words)
// 		fmt.Println(result)
// 	}
// }
// ascii-art-web
// Objectives

// Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.

// Your webpage must allow the use of the different banners:

//     shadow
//     standard
//     thinkertoy

// Implement the following HTTP endpoints:

//     GET /: Sends HTML response, the main page.
//     1.1. GET Tip: go templates to receive and display data from the server.

//     POST /ascii-art: that sends data (text and specified banner) to the Go server.
//     2.1. POST Tip: use form and other types of tags to make the post request.

// The way you display the result from the POST is up to you. What we recommend are one of the following :

//     Display the result in the route /ascii-art after the POST is completed. So going from the home page to another page.
//     Or display the result of the POST in the home page. This way appending the results in the home page.

// The main page must have:

//     a text input field
//     radio buttons, select object or another method to switch between banners
//     a button that sends a POST request to '/ascii-art' and displays the result on the page.

// HTTP status code

// Your endpoints must return appropriate HTTP status codes.

//     200 OK, if everything went without errors.
//     404 Not Found, if nothing is found, for example templates or banners.
//     400 Bad Request, for incorrect requests.
//     500 Internal Server Error, for unhandled errors.

// Markdown

// In the root project directory create a README.MD file with the following sections and contents:

//     Description
//     Authors
//     Usage: how to run
//     Implementation details: algorithm

// Instructions

//     HTTP server must be written in Go.
//     HTML templates must be in the project root directory templates.
//     The code must respect the good practices.

// Allowed packages

//     Only the standard go packages are allowed

// i just want the func main.go to be the main web server calling those funcs to work dont worry with my exiting code it works
