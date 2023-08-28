package main

import (
	"log"
	"os"

	"github.com/WanderningMaster/ssetup.git/internal/render"
	"github.com/WanderningMaster/ssetup.git/internal/store"
)

func main() {
	args := os.Args[1:]
	store.SetupStore()
	if len(args) == 0 {
		render.Loop()
	}

	if args[0] == "add" && len(args) == 2 {
		store.AddFile(args[1])
	} else {
		log.Fatal("invalid command")
	}
}
