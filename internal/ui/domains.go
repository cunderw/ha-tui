package ui

import (
	"github.com/cunderw/ha-tui/internal/data"
	"github.com/cunderw/ha-tui/internal/model"
	"github.com/rivo/tview"
)

type DomainPane struct {
	*tview.Flex
	list    *tview.List
	domains []model.Domain
}

func NewDomainsPane() *DomainPane {
	pane := DomainPane{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		list: tview.NewList().ShowSecondaryText(false),
	}

	pane.SetBorder(true).SetTitle("[::u]D[::-]omains")

	pane.AddItem(pane.list, 0, 1, true)

	return &pane
}

// ClearList removes all items from TaskPane
func (pane *DomainPane) ClearList() {
	pane.list.Clear()
	pane.domains = nil
}

// SetList Sets a list of tasks to be displayed
func (pane *DomainPane) SetList(domains []model.Domain) {
	pane.ClearList()
	pane.domains = domains

	for i := range pane.domains {
		pane.addDomainToList(i)
	}
}

func (pane *DomainPane) addDomainToList(i int) *tview.List {
	return pane.list.AddItem(pane.domains[i].Name, "", 0, func() {
		entities := data.GetEntities((&pane.domains[i].Name))
		entitiesPane.SetList(entities)

		if len(entities) > 0 {
			app.SetFocus(entitiesPane.list)
		}

	})
}
