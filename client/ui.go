package client

import (
	"strconv"

	tui "github.com/marcusolsson/tui-go"
)

var ui tui.UI

// RenderUI hi
func renderUI(ip string) {

	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)

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

	statusbar := tui.NewStatusBar("Hello this is test")
	statusbar.SetStyleName("warning")

	messagesbox := tui.NewVBox(
		// messages,
		statusbar,
		scrollbar,
	)
	messagesbox.SetSizePolicy(tui.Expanding, tui.Expanding)

	messagesbox.SetBorder(true)

	sd := tui.NewLabel("Scrollbar" + strconv.Itoa(scrollbar.Size().Y))
	msgd := tui.NewLabel("Messages" + strconv.Itoa(messages.Size().Y))
	msgbd := tui.NewLabel("Messagesbox" + strconv.Itoa(messagesbox.Size().Y))

	sidebar := tui.NewVBox(
		tui.NewLabel("Connected"),
		tui.NewButton(ip),
		tui.NewSpacer(),
		tui.NewLabel("-DEBUG-"),
		tui.NewSpacer(),
		sd,
		msgd,
		msgbd,
	)

	sidebar.SetBorder(true)

	sidebar.SetSizePolicy(tui.Maximum, tui.Minimum)

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

	ui.SetTheme(t)

	ui.SetKeybinding("Ctrl+C", func() { sendMsg(srv.Conn, "_CLOSE"); closeConn(srv.Conn); ui.Quit() })
	// ui.SetKeybinding("Up", func() { scrollbar.Scroll(0, 1) })
	ui.SetKeybinding("Up", func() {
		if messages.Selected() >= 0 {
			messages.Select(messages.Selected() - 1)
			if -(messages.Selected()-messages.Length())-1 == messages.Selected() && messages.Length() > scrollbar.Size().Y {
				scrollbar.Scroll(0, -1)
			}
		} else {
			messages.Select(messages.Length() - 1)
		}
		sd.SetText("Scrollbar" + strconv.Itoa(scrollbar.Size().Y))
		msgd.SetText("Messages" + strconv.Itoa(messages.Size().Y))
		msgbd.SetText("Messagesbox" + strconv.Itoa(messagesbox.Size().Y))
	})
	// ui.SetKeybinding("Down", func() { scrollbar.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() {
		if messages.Length() >= messages.Selected()+2 {
			messages.Select(messages.Selected() + 1)
			if messages.Selected() == messages.Length() && messages.Length() > scrollbar.Size().Y {
				scrollbar.Scroll(0, -1)
			}
		}
		sd.SetText("Scrollbar" + strconv.Itoa(scrollbar.Size().Y))
		msgd.SetText("Messages" + strconv.Itoa(messages.Size().Y))
		msgbd.SetText("Messagesbox" + strconv.Itoa(messagesbox.Size().Y))
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
