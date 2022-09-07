package main

import (
	"fmt"
	"gee/config"
)

// Main program.
func main() {
	conf := config.New()

	fmt.Printf("conf: %v\n", conf)
}
