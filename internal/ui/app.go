package ui

import (
	"unicode"

	"github.com/cunderw/ha-tui/internal/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app              *tview.Application
	layout, contents *tview.Flex

	titleBar     *TitleBar
	domainsPane  *DomainPane
	entitiesPane *EntitiesPane
	entityPane   *EntityPane
)

func StartApp() *tview.Application {
	app = tview.NewApplication()
	titleBar = NewTitleBar()
	domainsPane = NewDomainsPane()
	entitiesPane = NewEntitiesPane()
	entityPane = NewEntityPane()

	domains := data.GetDomains()

	domainsPane.SetList(domains)

	contents := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(domainsPane, 0, 1, false).
		AddItem(entitiesPane, 0, 1, false).
		AddItem(entityPane, 0, 3, false)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titleBar, 2, 1, false).
		AddItem(contents, 0, 1, false)

	app.SetRoot(layout, true).SetFocus(domainsPane)

	setKeyboardShortcuts()

	return app
}

func setKeyboardShortcuts() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// Global shortcuts
		switch unicode.ToLower(event.Rune()) {
		case 'd':
			app.SetFocus(domainsPane)
			return nil
		case 'e':
			app.SetFocus(entitiesPane)
			return nil
		}

		return event
	})
}
