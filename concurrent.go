package main

import (
	"fmt"
	"time"
)

func say(snippet string, delay time.Duration) {
	for {
		fmt.Println(snippet)
		time.Sleep(delay)
	}
}

func main() {
	go say("Hello World!", time.Second)
	go say("Yeah!", 5*time.Second)

	time.Sleep(10 * time.Second)
}
