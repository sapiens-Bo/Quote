package display

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Display(quote, author string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = quote
	p.TextStyle = ui.NewStyle(ui.ColorYellow)
	p.SetRect(0, 0, 50, 5)
	p.Title = author
	p.TitleStyle = ui.NewStyle(ui.ColorMagenta)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
