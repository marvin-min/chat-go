package main

import (
	"fmt"
	"time"
)

func main() {
	server := NewServer("0.0.0.0", 9999)
	// go alert()
	server.Start()
}

func alert() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("你该休息了")
	}
}
