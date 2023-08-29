package commands

import (
	"log"

	"github.com/WanderningMaster/ssetup.git/internal/store"
)

type CommandIdent = string

const (
	Add CommandIdent = "add"
)

type Command struct {
	ident CommandIdent
	exec  func([]string)
}

func addCommand(args []string) {
	if args[0] == "add" && len(args) == 2 {
		store.AddFile(args[1])
	} else {
		log.Fatal("invalid command")
	}
}

func ProcessArgs(args []string) {
	cmd := map[CommandIdent]func([]string){
		Add: addCommand,
	}

	cmd[args[0]](args)
}
