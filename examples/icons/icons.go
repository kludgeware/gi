// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"

	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
)

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 768
	nColumns := 5

	oswin.TheApp.SetName("icons")
	oswin.TheApp.SetAbout(`This is a demo of the icons in the <b>GoGi</b> graphical interface system, within the <b>GoKi</b> tree framework.  See <a href="https://github.com/goki">GoKi on GitHub</a>`)

	win := gi.NewWindow2D("gogi-icons-demo", "GoGi Icons", width, height, true)

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()

	row1 := mfr.AddNewChild(gi.KiT_Layout, "row1").(*gi.Layout)
	row1.Lay = gi.LayoutHoriz
	row1.SetProp("margin", 2.0) // raw numbers = px = 96 dpi pixels
	row1.SetStretchMaxWidth()

	spc := mfr.AddNewChild(gi.KiT_Space, "spc1").(*gi.Space)
	spc.SetFixedHeight(units.NewValue(2.0, units.Em))

	row1.AddNewChild(gi.KiT_Stretch, "str1")
	lab1 := row1.AddNewChild(gi.KiT_Label, "lab1").(*gi.Label)
	lab1.Text = "These are all the GoGi Icons, in a small and large size"
	lab1.SetProp("max-width", -1)
	lab1.SetProp("text-align", "center")
	row1.AddNewChild(gi.KiT_Stretch, "str2")

	grid := mfr.AddNewChild(gi.KiT_Frame, "grid").(*gi.Frame)
	grid.Lay = gi.LayoutGrid
	grid.Stripes = gi.RowStripes
	grid.SetProp("columns", nColumns)
	grid.SetProp("horizontal-align", "center")
	grid.SetProp("margin", 2.0)
	grid.SetProp("spacing", 6.0)
	grid.SetStretchMaxWidth()
	grid.SetStretchMaxHeight()

	il := gi.TheIconMgr.IconList(true)

	for _, icnm := range il {
		icnms := string(icnm)
		if icnms == "none" {
			continue
		}
		vb := grid.AddNewChild(gi.KiT_Layout, "vb").(*gi.Layout)
		vb.Lay = gi.LayoutVert

		nmlbl := vb.AddNewChild(gi.KiT_Label, "lab1").(*gi.Label)
		nmlbl.Text = icnms

		smico := vb.AddNewChild(gi.KiT_Icon, icnms).(*gi.Icon)
		smico.SetIcon(icnms)
		smico.SetMinPrefWidth(units.NewValue(20, units.Px))
		smico.SetMinPrefHeight(units.NewValue(20, units.Px))
		smico.SetProp("background-color", color.Transparent)
		smico.SetProp("fill", "#88F")
		smico.SetProp("stroke", "black")
		// smico.SetProp("horizontal-align", gi.AlignLeft)

		ico := vb.AddNewChild(gi.KiT_Icon, icnms).(*gi.Icon)
		ico.SetIcon(icnms)
		ico.SetMinPrefWidth(units.NewValue(100, units.Px))
		ico.SetMinPrefHeight(units.NewValue(100, units.Px))
		ico.SetProp("background-color", color.Transparent)
		ico.SetProp("fill", "#88F")
		ico.SetProp("stroke", "black")
		// ico.SetProp("horizontal-align", gi.AlignLeft)
	}

	// main menu
	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.KnownChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.KnownChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.OSWin.SetCloseCleanFunc(func(w oswin.Window) {
		go oswin.TheApp.Quit() // once main window is closed, quit
	})

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()
}
