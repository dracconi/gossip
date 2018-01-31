package client

import (
	tui "github.com/marcusolsson/tui-go"
)

// RenderUI hi
func renderUI(ip string) {

	box := tui.NewVBox()

	box.SetBorder(true)
	box.SetTitle(ip)

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

	messagesbox := tui.NewHBox(
		// messages,
		scrollbar,
	)
	messagesbox.SetSizePolicy(tui.Expanding, tui.Expanding)

	messagesbox.SetBorder(true)

	wrapper := tui.NewHBox(
		box,
		tui.NewVBox(
			messagesbox,
			inputbox,
		),
	)

	uinput.OnSubmit(func(e *tui.Entry) {
		sendMsg(srv.Conn, e.Text())
		// e.SetText("")
	})

	go fetchMsg(srv.Conn, messages, scrollbar)

	ui, err := tui.New(wrapper)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { sendMsg(srv.Conn, "_CLOSE\n"); closeConn(srv.Conn); ui.Quit() })
	ui.SetKeybinding("Up", func() { scrollbar.Scroll(0, 1) })
	ui.SetKeybinding("Down", func() { scrollbar.Scroll(0, -1) })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
