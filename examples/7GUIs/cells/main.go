package main

import (
	"examples/7GUIs/cells/seven_gui_cells_panel"
	"examples/7GUIs/cells/sven_gui_cells_context"
	"github.com/george012/fltk_go"
)

const (
	WIDGET_HEIGHT = 200
	WIDGET_WIDTH  = 450
)

const (
	MAX_ROW_COUNT = 100 // 0~99
	MAX_COL_COUNT = 26  // A~Z
)

var ctx = sven_gui_cells_context.NewContext(MAX_ROW_COUNT, MAX_COL_COUNT)

func init() {
	ctx.UpdateCellAtLoc("B1", "5")
	ctx.UpdateCellAtLoc("B2", "1")
	ctx.UpdateCellAtLoc("B3", "10.3")
	ctx.UpdateCellAtLoc("B4", "22.87")
	ctx.UpdateCellAtLoc("B5", "=SUM(B1:B4)")
	ctx.UpdateCellAtLoc("C1", "6")
	ctx.UpdateCellAtLoc("C2", "7")
	ctx.UpdateCellAtLoc("C3", "2")
	ctx.UpdateCellAtLoc("C4", "5")
	ctx.UpdateCellAtLoc("C5", "=SUM(C1:C4)")
	ctx.UpdateCellAtLoc("A5", "Sum")
	ctx.UpdateCellAtLoc("D5", "=SUM(B5:C5)")
}

func main() {
	fltk_go.SetScheme("gtk+")

	win := fltk_go.NewWindow(
		WIDGET_WIDTH,
		WIDGET_HEIGHT)
	win.SetLabel("Cells")

	p := seven_gui_cells_panel.NewPanel(win, MAX_ROW_COUNT, MAX_COL_COUNT, ctx)
	p.Bind(ctx)

	win.End()
	win.Show()
	fltk_go.Run()
}
