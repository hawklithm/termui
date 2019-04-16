// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"

	ui "github.com/hawklithm/termui"
	"github.com/hawklithm/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewImageList()
	l.Title = "List"
	p := []string{
		"[0] github.com/hawklithm/termui",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[5] random_out.go\n12dfasdf\n1fldjalgd",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}
	z := []*widgets.ImageListItem{}
	for _, s := range p {
		item := widgets.NewImageListItem()
		item.Text = s
		z = append(z, item)
	}
	l.Rows = z
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 24)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
