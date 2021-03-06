// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
)

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 768

	win := gi.NewWindow2D("gogi-tabview-test", "GoGi TabView Test", width, height, true) // pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()

	tv := mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)
	tv.NewTabButton = true

	lbl1k, _ := tv.AddNewTab(gi.KiT_Label, "This is Label1")
	lbl1 := lbl1k.(*gi.Label)
	lbl1.SetText("this is the contents of the first tab")
	lbl1.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

	lbl2k, _ := tv.AddNewTab(gi.KiT_Label, "And this Label2")
	lbl2 := lbl2k.(*gi.Label)
	lbl2.SetText("this is the contents of the second tab")
	lbl2.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

	tv.SelectTabIndex(0)

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
