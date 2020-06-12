package gui

import (
	"log"

	"github.com/asaskevich/EventBus"
	"github.com/awesome-gocui/gocui"
)

// InitGUI - Initialise the gui
func InitGUI(bus EventBus.Bus) {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
