package main

import (
	"fmt"
	"strings"
)

func gameOver() {
	fmt.Println("Game over. You can try again.")
}

func invalidWarning() {
	fmt.Println("Error: The answer you provided is invalid. Please select A, B, C, or D.")
}

func gameWon(username string, totalAmount int) {
	if strings.TrimSpace(username) == "" {
		username = "anonymous"
	}
	fmt.Printf("Congratulations %s, you successfully won $%d, contact the developers of this site to receive payment.\n", username, totalAmount)
}

func cleanInput(input string) string {
	return strings.TrimSpace(strings.ToUpper(input))
}

func validChoice(choice string) bool {
	return choice == "A" || choice == "B" || choice == "C" || choice == "D"
}

func calculatePercentage(correct, total int) string {
	if total == 0 {
		return "0.0%"
	}
	percent := (float64(correct) / float64(total)) * 100
	return fmt.Sprintf("%.1f%%", percent)
}

func status(level string, category string, currentWallet int) {
	fmt.Printf("[%s - Topic: %s] Current Balance: $%d\n", level, category, currentWallet)
}
