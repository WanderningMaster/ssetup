package render

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/WanderningMaster/ssetup.git/internal/store"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter script name: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	dir := store.GetTempDir()

	cmd := exec.Command("vim", dir+"/"+input)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func Loop() {
	manageMenu := NewMenu("Select option:")
	manageMenu.AddItem("New script", string(NewScript), openNewScript)
	manageMenu.AddItem("Edit script", string(EditScript), nil)
	manageMenu.AddItem("Delete script", string(DeleleScript), nil)

	rootMenu := NewMenu("Select option:")
	rootMenu.AddItem("Run", string(Run), nil)
	rootMenu.AddItemWithSubMenu("Manage scripts", string(Manage), manageMenu)

	for {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}
		cmd := rootMenu.Loop()
		if cmd == nil {
			break
		}
		keyboard.Close()
		cmd()
	}
}
