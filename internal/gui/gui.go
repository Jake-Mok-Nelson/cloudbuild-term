package gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/jroimartin/gocui"
)

func NextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "projects" {
		_, err := g.SetCurrentView("builds")
		return err
	}
	_, err := g.SetCurrentView("projects")
	return err
}

func CursorDown(g *gocui.Gui, v *gocui.View) error {
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

func CursorUp(g *gocui.Gui, v *gocui.View) error {
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

func GetLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func DelMsg(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("projects"); err != nil {
		return err
	}
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("projects", gocui.KeyCtrlSpace, gocui.ModNone, NextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("builds", gocui.KeyCtrlSpace, gocui.ModNone, NextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyArrowDown, gocui.ModNone, CursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyArrowUp, gocui.ModNone, CursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyEnter, gocui.ModNone, GetLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, DelMsg); err != nil {
		return err
	}

	if err := g.SetKeybinding("builds", gocui.KeyCtrlS, gocui.ModNone, Savebuilds); err != nil {
		return err
	}
	if err := g.SetKeybinding("builds", gocui.KeyCtrlW, gocui.ModNone, SaveVisualbuilds); err != nil {
		return err
	}
	return nil
}

func Savebuilds(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveVisualbuilds(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// HEADER VIEW
	if v, err := g.SetView("header", -1, -1, maxX, 30); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = true
		fmt.Fprintln(v, `
   ___ _                 _ _           _ _     _   _____
  / __\ | ___  _   _  __| | |__  _   _(_) | __| | /__   \___ _ __ _ __ ___
 / /  | |/ _ \| | | |/ _' | '_ \| | | | | |/ _' |   / /\/ _ \ '__| '_ ' _ \
/ /___| | (_) | |_| | (_| | |_) | |_| | | | (_| |  / / |  __/ |  | | | | | |
\____/|_|\___/ \__,_|\__,_|_.__/ \__,_|_|_|\__,_|  \/   \___|_|  |_| |_| |_|
  `)

		fmt.Fprint(v, "\n\n")
		if _, err := g.SetCurrentView("header"); err != nil {
			return err
		}
	}

	// PROJECTS VIEW
	if v, err := g.SetView("projects", 0, 10, 30, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Projects"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		fmt.Fprintln(v, "All")
		fmt.Fprintln(v, "SomeGCPProject")
		fmt.Fprintln(v, "AnotherGCPproject")
		fmt.Fprintln(v, "YetAnotherProject")

	}

	// BUILD STATUS VIEW
	if v, err := g.SetView("builds", 31, 10, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Build Status"
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("builds"); err != nil {
			return err
		}
	}
	return nil
}
