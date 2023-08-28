package render

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/WanderningMaster/ssetup.git/internal/store"
	"github.com/WanderningMaster/ssetup.git/internal/utils"
	"github.com/eiannone/keyboard"
)

type RootMenuKeys string
type ManageMenuKeys string

const (
	Run    RootMenuKeys = "Run"
	Manage RootMenuKeys = "Manage"
	Exit   RootMenuKeys = "Exit"
)
const (
	NewScript    ManageMenuKeys = "New script"
	EditScript   ManageMenuKeys = "Edit script"
	DeleleScript ManageMenuKeys = "Delete script"
	Back         ManageMenuKeys = "Back"
)

func openNewScript() {
	keyboard.Close()
	fmt.Print("Enter script name: ")

	input := utils.ReadLine()
	input = strings.TrimRight(input, "\r\n")
	dir := store.GetLocalDataDir()

	cmd := exec.Command("vim", dir+"/"+input)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
}

func runScript(sc store.Script) {
	keyboard.Close()
	cmd := exec.Command(sc.Exec, sc.Path)

	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	wg.Add(1)

	go Loader(&wg, stopChan)
	cmd.Run()

	close(stopChan)
	wg.Wait()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
}

func RenderScriptList() {
	list, err := store.ListScripts()
	if err != nil {
		log.Fatal(err.Error())
	}
	menu := NewMenu("")
	for _, sc := range list {
		scCopy := sc
		cmd := func() {
			runScript(scCopy)
		}

		menu.AddItem(sc.Name, sc.Path, cmd)
	}

	menu.Loop()
}

func Loop() {
	manageMenu := NewMenu("Select option:")
	manageMenu.AddItem("New script", string(NewScript), openNewScript)
	manageMenu.AddItem("Edit script", string(EditScript), nil)
	manageMenu.AddItem("Delete script", string(DeleleScript), nil)

	rootMenu := NewMenu("Select option:")
	rootMenu.AddItem("Run", string(Run), RenderScriptList)
	rootMenu.AddItemWithSubMenu("Manage scripts", string(Manage), manageMenu)

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	rootMenu.Loop()
}
