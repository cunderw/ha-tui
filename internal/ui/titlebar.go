package ui

import "github.com/rivo/tview"

type TitleBar struct {
	*tview.Flex
	titleText   *tview.TextView
	versionText *tview.TextView
}

func NewTitleBar() *TitleBar {
	bar := TitleBar{
		Flex:        tview.NewFlex().SetDirection(tview.FlexRow),
		titleText:   tview.NewTextView().SetText("[lime::b]HA-TUI [::-]- Home Assistant In The Terminal").SetDynamicColors(true),
		versionText: tview.NewTextView().SetText("[::d]Version: 0.1.0").SetTextAlign(tview.AlignRight).SetDynamicColors(true),
	}

	bar.
		AddItem(bar.titleText, 0, 2, false).
		AddItem(bar.versionText, 0, 1, false)

	return &bar
}
