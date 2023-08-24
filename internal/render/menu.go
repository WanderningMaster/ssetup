package render

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

type Menu struct {
	Prompt    string
	CursorPos int
	Items     []*MenuItem
}

type MenuItem struct {
	Text    string
	ID      string
	SubMenu *Menu
	Command func()
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:    prompt,
		CursorPos: 0,
		Items:     make([]*MenuItem, 0),
	}
}

func (m *Menu) AddItemWithSubMenu(text string, id string, subMenu *Menu) {
	m.Items = append(m.Items, &MenuItem{
		Text:    text,
		ID:      id,
		SubMenu: subMenu,
		Command: nil,
	})
}

func (m *Menu) AddItem(text string, id string, command func()) {
	m.Items = append(m.Items, &MenuItem{
		Text:    text,
		ID:      id,
		SubMenu: nil,
		Command: command,
	})
}

func (m *Menu) Render() {
	for i, item := range m.Items {
		if i == m.CursorPos {
			fmt.Printf("> %s\n", item.Text)
		} else {
			fmt.Printf("  %s\n", item.Text)
		}
	}
}

func (m *Menu) MoveCursorUp() func() {
	if m.CursorPos > 0 {
		m.CursorPos -= 1
	}

	return nil
}

func (m *Menu) MoveCursorDown() func() {
	if m.CursorPos < len(m.Items)-1 {
		m.CursorPos += 1
	}

	return nil
}

func (m *Menu) Select() func() {
	item := m.Items[m.CursorPos]
	if item.SubMenu != nil {
		return item.SubMenu.Loop()
	}
	return item.Command
}

func (m *Menu) runCommand(key keyboard.Key) func() {
	var cmd = map[keyboard.Key]func() func(){
		keyboard.KeyArrowUp:   m.MoveCursorUp,
		keyboard.KeyArrowDown: m.MoveCursorDown,
		keyboard.KeyEnter:     m.Select,
	}
	return cmd[key]()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (m *Menu) Loop() func() {
	for {
		clearScreen()
		m.Render()
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
		}

		if char == 'q' || key == keyboard.KeyEsc {
			clearScreen()
			break
		}
		cmd := m.runCommand(key)
		if cmd != nil {
			return cmd
		}
	}

	return nil
}
