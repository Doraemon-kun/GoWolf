package backend

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func getTerminalSize() (int, int) {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return width, height
}

func writeLineBreak() {
	w, _ := getTerminalSize()
	fmt.Println(strings.Repeat("-", w))
}

func options() {
	for {
		var ch string
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. New Game")
		fmt.Println("2. Help")
		fmt.Println("3. Quit")
		fmt.Printf("> ")
		fmt.Scanln(&ch)
		if ch_i, err := strconv.ParseInt(ch, 10, 8); err == nil {
			if 0 < ch_i && ch_i <= 3 {
				fmt.Println("You choose option", fmt.Sprintf("%d", ch_i)) // Placeholder
				break
			} else {
				RaiseError(2)
				writeLineBreak()
				continue
			}
		} else {
			RaiseError(2)
			writeLineBreak()
			continue
		}
	}
}

func Intro() {
	fmt.Println("Welcome to GoWolf! The Werewolf game for Terminal users!")
	fmt.Println("Version 0.1a")
	writeLineBreak()
	options()
}
