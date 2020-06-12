package gui

import (
	"encoding/json"
	"fmt"

	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/builds"
	"github.com/jroimartin/gocui"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "projects" {
		_, err := g.SetCurrentView("builds")
		return err
	}
	_, err := g.SetCurrentView("projects")
	return err
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// Quit - Quit our application
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// SelectProject - Select a project from the project view
func selectProject(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, oy := v.Cursor()
		lineText, err := v.Line(oy)
		if err != nil {
			return err
		}

		// Someone selected a project, we will populate the builds view based on the selection
		updateBuildsListView(g, lineText)

	}
	return nil
}

// list the builds in the builds view
func updateBuildsListView(g *gocui.Gui, project string) error {

	view, err := g.SetCurrentView("builds")
	if err != nil {
		return err
	}

	// Fetch a list of builds for a given projectId
	buildsData, err := builds.FetchBuilds(project, 2)
	if err != nil {
		return err
	}

	// Update the builds view with the output of the builds command
	view.Clear()

	// // Creating the maps for JSON
	// m := map[string]interface{}{}
	b := builds.BuildOverview{}

	// // Parsing/Unmarshalling JSON encoding/json
	err = json.Unmarshal([]byte(buildsData), &b)
	if err != nil {
		return err
	}

	for _, val := range b.Builds {
		fmt.Printf(val.BuildTriggerID, view)

	}

	view.Write(buildsData)
	// TODO: Print out the list of builds here

	return nil
}
