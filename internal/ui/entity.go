package ui

import (
	"github.com/cunderw/ha-tui/internal/model"
	"github.com/rivo/tview"
)

type EntityPane struct {
	*tview.Flex
	entity model.Entity
}

func NewEntityPane() *EntityPane {
	pane := EntityPane{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
	}

	pane.SetBorder(true).SetTitle("[::u]D[::-]etails")

	return &pane
}

func (pane *EntityPane) SetDetails(entity model.Entity) {
	pane.entity = entity
	pane.Clear()

	icon := tview.NewTextView().SetText(GetIcon(pane.entity.Domain))
	name := tview.NewTextView().SetText(pane.entity.Name)

	pane.AddItem(icon, 0, 1, false)
	pane.AddItem(name, 0, 1, false)
}
