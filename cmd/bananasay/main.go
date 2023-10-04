package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// nolint forbidigo
func main() {
	var text string
	if len(os.Args) > 1 {
		text = strings.Join(os.Args[1:], " ")
	} else {
		text = getStdin()
	}
	fmt.Println(bubble(text))
	fmt.Println(banana())
}

func getStdin() string {
	var text string
	input := make(chan string)
	go func(ch chan string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		} else {
			ch <- text
		}
	}(input)
	select {
	case input := <-input:
		text = input
	case <-time.After(time.Second):
		text = quote()
	}
	return text
}
