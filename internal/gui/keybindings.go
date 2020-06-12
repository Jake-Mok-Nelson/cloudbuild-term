package gui

import "github.com/awesome-gocui/gocui"

// Keybindings - Keybindings for views in our layout
func Keybindings(g *gocui.Gui) error {

	// Keybindings for all views

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Keybindings for the projects view

	if err := g.SetKeybinding("projects", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := g.SetKeybinding("projects", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("projects", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("projects", gocui.KeyEnter, gocui.ModNone, selectProject); err != nil {
		return err
	}

	// Keybindings for the builds view

	if err := g.SetKeybinding("builds", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := g.SetKeybinding("builds", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("builds", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	return nil
}
