package client

import (
	"strconv"

	tui "github.com/marcusolsson/tui-go"
)

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

	scrollbar := &scrollArea{tui.NewScrollArea(messages), messages, 0}

	t.SetStyle("label.statusbar", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})

	messagesbox := tui.NewVBox(
		// messages,
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
	ui.SetKeybinding("Up", func() { scrollbar.Scroll(0, -1) })
	// ui.SetKeybinding("Up", func() {
	// 	if messages.Selected() >= 0 {
	// 		messages.Select(messages.Selected() - 1)
	// 		if -(messages.Selected()-messages.Length())-1 == messages.Selected() && messages.Length() > scrollbar.Size().Y {
	// 			scrollbar.Scroll(0, -1)
	// 		}
	// 	} else {
	// 		messages.Select(messages.Length() - 1)
	// 	}
	// 	sd.SetText("Scrollbar" + strconv.Itoa(scrollbar.Size().Y))
	// 	msgd.SetText("Messages" + strconv.Itoa(messages.Size().Y))
	// 	msgbd.SetText("Messagesbox" + strconv.Itoa(messagesbox.Size().Y))
	// })
	ui.SetKeybinding("Down", func() { scrollbar.Scroll(0, 1) })
	// ui.SetKeybinding("Down", func() {
	// 	if messages.Length() >= messages.Selected()+2 {
	// 		messages.Select(messages.Selected() + 1)
	// 		if messages.Selected() == messages.Length() && messages.Length() > scrollbar.Size().Y {
	// 			scrollbar.Scroll(0, -1)
	// 		}
	// 	}
	// 	sd.SetText("Scrollbar" + strconv.Itoa(scrollbar.Size().Y))
	// 	msgd.SetText("Messages" + strconv.Itoa(messages.Size().Y))
	// 	msgbd.SetText("Messagesbox" + strconv.Itoa(messagesbox.Size().Y))
	// })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
