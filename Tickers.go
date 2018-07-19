package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(16000 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
