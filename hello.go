package main

import (
	"fmt"
	"time"
)

const (
	cwidSuffix = "1111" // Replace with your last four CWID digits
	timeFormat = "2006-01-02 15:04:05"
)

func main() {
	currentTime := time.Now().Format(timeFormat)
	fmt.Println("Hello, world!")
	fmt.Printf("It's %s now, and my CWID ending in %s.\n", currentTime, cwidSuffix)
}
