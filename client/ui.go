package client

import (
	tui "github.com/marcusolsson/tui-go"
)

// type messagesList struct {
// 	*tui.Table
// }

// func (tab *messagesList) appendMsg(msg string) {
// 	tab.AppendRow(
// 		tui.NewHBox(
// 			tui.NewLabel(msg),
// 		),
// 	)
// }

type selectWidget interface {
	tui.Widget
	Select(i int)
	Selected() int
}

type scrollArea struct {
	*tui.ScrollArea
	wrapped selectWidget
	cur     int
}

func (s *scrollArea) Scroll(_, y int) {
	selected := s.wrapped.Selected()
	scrollTo := 0

	defer func() {
		s.wrapped.Select(selected)
		s.ScrollArea.Scroll(0, scrollTo)
	}()

	// UPWARDS
	if y == -1 {

		// if row is not top, select one upper
		if selected != 0 {
			selected--
			if selected > s.cur {
				s.cur--
				scrollTo--
			}
			return

		}

		if s.cur == 0 {
			return
		}
		return
	}

	// DOWNWARDS

	if selected != s.wrapped.Size().Y-1 {
		selected++

		if selected > s.cur+s.ScrollArea.Size().Y-1 {
			s.cur++
			scrollTo++
		}
		return
	}

}

var ui tui.UI

// RenderUI hi
func renderUI(ip string) {

	uinput := tui.NewEntry()
	uinput.SetFocused(true)
	inputbox := tui.NewHBox(uinput)
	inputbox.SetBorder(true)

	uinput.SetSizePolicy(tui.Expanding, tui.Maximum)
	inputbox.SetSizePolicy(tui.Expanding, tui.Maximum)

	messages := tui.NewList()

	messages.SetSizePolicy(tui.Expanding, tui.Maximum)

	scrollbar := &scrollArea{tui.NewScrollArea(messages), messages, 1}
	messages.Select(1)

	messagesbox := tui.NewVBox(
		// messages,
		scrollbar,
	)
	messagesbox.SetSizePolicy(tui.Expanding, tui.Expanding)

	messagesbox.SetBorder(true)

	sidebar := tui.NewVBox(
		tui.NewLabel("Connected"),
		tui.NewButton(ip),
		tui.NewSpacer(),
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

	ui, err := tui.New(wrapper)
	if err != nil {
		panic(err)
	}

	go fetchMsg(srv.Conn, messages, scrollbar, ui)
	ui.SetKeybinding("Ctrl+D", func() { sendMsg(srv.Conn, "_CLOSE"); closeConn(srv.Conn); ui.Quit() })
	ui.SetKeybinding("Up", func() { scrollbar.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() { scrollbar.Scroll(0, 1) })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
