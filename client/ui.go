package client

import (
	tui "github.com/marcusolsson/tui-go"
)

// RenderUI hi
func renderUI(ip string) {

	sidebar := tui.NewVBox(
		tui.NewLabel("Connected"),
		tui.NewButton(ip),
		tui.NewSpacer(),
	)

	sidebar.SetBorder(true)

	sidebar.SetSizePolicy(tui.Maximum, tui.Minimum)

	uinput := tui.NewEntry()
	uinput.SetFocused(true)
	inputbox := tui.NewHBox(uinput)
	inputbox.SetBorder(true)

	uinput.SetSizePolicy(tui.Expanding, tui.Maximum)
	inputbox.SetSizePolicy(tui.Expanding, tui.Maximum)

	messages := tui.NewList()
	messages.AddItems("~ CONNECTED TO " + ip + " ~")

	messages.SetSizePolicy(tui.Expanding, tui.Maximum)

	scrollbar := tui.NewScrollArea(messages)

	messagesbox := tui.NewVBox(
		// messages,
		scrollbar,
	)
	messagesbox.SetSizePolicy(tui.Expanding, tui.Expanding)

	messagesbox.SetBorder(true)

	wrapper := tui.NewHBox(
		sidebar,
		tui.NewVBox(
			messagesbox,
			inputbox,
		),
	)

	uinput.OnSubmit(func(e *tui.Entry) {
		sendMsg(srv.Conn, e.Text())
	})

	go fetchMsg(srv.Conn, messages, scrollbar)

	ui, err := tui.New(wrapper)
	if err != nil {
		panic(err)
	}

	ui.SetKeybinding("Esc", func() { sendMsg(srv.Conn, "_CLOSE"); closeConn(srv.Conn); ui.Quit() })
	// ui.SetKeybinding("Up", func() { scrollbar.Scroll(0, 1) })
	ui.SetKeybinding("Up", func() {
		messages.SetSelected(messages.Selected() - 1)

	})
	// ui.SetKeybinding("Down", func() { scrollbar.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() {
		if messages.Length() >= messages.Selected()+2 {
			messages.Select(messages.Selected() + 1)
			if messages.Selected() == messages.Length() && messages.Length() > scrollbar.Size().Y {
				scrollbar.Scroll(0, -1)
			}
		}
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
