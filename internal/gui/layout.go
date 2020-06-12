package gui

import (
	"fmt"

	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/projects"
	"github.com/awesome-gocui/gocui"
)

// Layout - The viusal layout of our application
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// HEADER VIEW
	if v, err := g.SetView("header", -1, -1, maxX, 90, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Editable = false
		v.FgColor = gocui.ColorBlue

		v.Wrap = true

		fmt.Fprint(v, `
     ___ _                 _ _           _ _     _   _____
    / __\ | ___  _   _  __| | |__  _   _(_) | __| | /__   \___ _ __ _ __ ___
   / /  | |/ _ \| | | |/ _' | '_ \| | | | | |/ _' |   / /\/ _ \ '__| '_ ' _ \
  / /___| | (_) | |_| | (_| | |_) | |_| | | | (_| |  / / |  __/ |  | | | | | |
  \____/|_|\___/ \__,_|\__,_|_.__/ \__,_|_|_|\__,_|  \/   \___|_|  |_| |_| |_|
  `)

		fmt.Fprint(v, "\n")
		fmt.Fprintln(v, "KEYBINDINGS")
		fmt.Fprintln(v, "Tab: Switch View         Enter: Select")
		fmt.Fprintln(v, "↑↓: Change Selection     ^C: Exit")
		fmt.Fprint(v, "\n")

		if _, err := g.SetCurrentView("header"); err != nil {
			return err
		}

	}

	// PROJECTS VIEW
	if v, err := g.SetView("projects", 0, 10, 30, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Projects"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorWhite

		p, err := projects.FetchProjects()
		if err != nil {
			//fmt.Errorf("unable to retrieve projects from Google: %v", err)
			panic(err)
		}
		for _, projectID := range p {
			fmt.Fprintln(v, projectID)
		}
	}

	// BUILD STATUS VIEW
	if v, err := g.SetView("builds", 31, 10, maxX-1, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Build Status"
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorWhite

		v.Wrap = true
		if _, err := g.SetCurrentView("builds"); err != nil {
			return err
		}

		// If project
	}
	return nil
}
