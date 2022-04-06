package ui

import (
	"github.com/cunderw/ha-tui/internal/model"
	"github.com/rivo/tview"
)

type EntitiesPane struct {
	*tview.Flex
	list     *tview.List
	entities []model.Entity
}

func NewEntitiesPane() *EntitiesPane {
	pane := EntitiesPane{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		list: tview.NewList().ShowSecondaryText(false),
	}

	pane.SetBorder(true).SetTitle("[::u]E[::-]ntities")

	pane.AddItem(pane.list, 0, 1, true)

	return &pane
}

// ClearList removes all items from TaskPane
func (pane *EntitiesPane) ClearList() {
	pane.list.Clear()
	pane.entities = nil
}

// SetList Sets a list of tasks to be displayed
func (pane *EntitiesPane) SetList(entities []model.Entity) {
	pane.ClearList()
	pane.entities = entities

	for i := range pane.entities {
		pane.addEntityToList(i)
	}
}

func (pane *EntitiesPane) addEntityToList(i int) *tview.List {
	return pane.list.AddItem(pane.entities[i].Name, "", 0, func() {
		entityPane.SetDetails(pane.entities[i])
	})
}
