package main

import (
	"math/rand"
	"strings"
)

var quotes = []string{
	"I'm waiting...",
	"What is yellow in color and can go 30 miles an hour?\n\nA banana inside a washing machine.",
	"What caused the baby banana to become so spoiled?\n\nHis mama left him out in the sun for too long.",
	"Do you know what I say to my bananas before leaving the house?\n\nI say, “I am going, bananas.”",
	"Another weekend has passed",
	"What do bananas say when they answer the phone?\n\nYellow.",
	"What's the best way to communicate with a banana?\n\nBy using a banana phone!",
	"Why don't most bananas tell secrets?\n\nBecause they're a fruit and lack the cognitive ability to communicate",
	"What's yellow and goes well in a smoothie?\n\nA banana. Because it's a banana.",
	"A banana's favorite place to visit is the peel-ippines.",
}

func quote() string {
	return quotes[rand.Intn(len(quotes))]
}

func bubble(text string) string {
	var (
		out   string
		arrow string
	)
	lines := strings.Split(text, "\n")
	maxWidth := 0
	for _, line := range lines {
		lineWidth := len(line)
		if lineWidth > maxWidth {
			maxWidth = lineWidth
		}
	}
	maxWidth += 2
	if maxWidth > 24 {
		arrow = "|/\n"
	} else {
		arrow = "\\|\n"
	}
	border := strings.Repeat("-", maxWidth)
	padding := strings.Repeat(" ", maxWidth/2)
	out += border + "\n"
	for _, line := range lines {
		out += "|" + line + strings.Repeat(" ", maxWidth-len(line)-1) + "|\n"
	}
	out += border + "\n" + padding + arrow
	return out
}
