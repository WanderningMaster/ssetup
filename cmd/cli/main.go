package main

import (
	"os"

	"github.com/WanderningMaster/ssetup.git/internal/commands"
	"github.com/WanderningMaster/ssetup.git/internal/render"
	"github.com/WanderningMaster/ssetup.git/internal/store"
)

func main() {
	args := os.Args[1:]
	store.SetupStore()
	if len(args) == 0 {
		render.Loop()
	} else {
		commands.ProcessArgs(args)
	}
}
