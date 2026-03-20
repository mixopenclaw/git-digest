package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("help", false, "Show help")
	version := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *help {
		fmt.Println("git-digest CLI\n\nUsage: git-digest [flags]\n\nFlags:\n  --help     Show help\n  --version  Show version")
		os.Exit(0)
	} else if *version {
		fmt.Println("git-digest version 0.1.0")
		os.Exit(0)
	}
}
