package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	com := &commands{}

	rand.Seed(time.Now().UnixNano())

	switch args[0] {
	case "gen":
		com.Generate()
	case "cp":
		com.Copy()
	default:
		flag.PrintDefaults()
	}
}
