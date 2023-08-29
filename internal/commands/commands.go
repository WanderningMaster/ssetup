package commands

import (
	"log"
	"sync"

	"github.com/WanderningMaster/ssetup.git/internal/render"
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

func execWithLoader(cmd func()) {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	wg.Add(1)

	go render.Loader(&wg, stopChan)
	cmd()

	close(stopChan)
	wg.Wait()
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
	execWithLoader(func() {
		cmd[args[0]](args)
	})
}
