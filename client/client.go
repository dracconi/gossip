package client

import (
	"bufio"
	"fmt"
	"net"

	"github.com/dracconi/gossip/handler"
	tui "github.com/marcusolsson/tui-go"
)

var srv handler.Server

func closeConn(conn net.Conn) {
	conn.Close()
}

// Loop that fetches messages
func fetchMsg(conn net.Conn, msgs *tui.List, scroll *scrollArea, ui tui.UI) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		msgs.AddItems(msg)
		ui.Update(func() {})
		scroll.Scroll(0, -1)
	}
}

// SendMsg Function that writes messages to the server
func sendMsg(conn net.Conn, msg string) {
	conn.Write([]byte(msg + "\n"))
}

// Client compartment
func Client(ipaddr string) {
	fmt.Println("CLIENT STARTED")
	srv.Conn, _ = net.Dial("tcp", ipaddr)
	if srv.Conn != nil {
		renderUI(ipaddr)
	} else {
		fmt.Println("Specified IP is not available.")
	}
}
