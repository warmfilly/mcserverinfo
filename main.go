package main

import (
	"fmt"
	"os"

	"github.com/FragLand/minestat/Go/minestat"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	server := args[0]

	fmt.Printf("%s\n", server)

	minestat.Init(server)

	if minestat.Online {
		displayServerInfo()
	} else {
		fmt.Println("Server is offline.")
	}
}

func displayServerInfo() {
	fmt.Println("Display info...")
	fmt.Printf("%d players online\n", minestat.Current_players)
}

func usage() {
	fmt.Println("Usage")
}
