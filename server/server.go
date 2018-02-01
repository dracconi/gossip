package server

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var conns []net.Conn

func broadcast(msg string) {
	for _, conn := range conns {
		conn.Write([]byte(msg))
	}
}

func handleConnection(conn net.Conn) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		if msg == "_CLOSE\n" {
			// for i, v := range conns {
			// 	if v == conn {
			// 		conns = append(conns[:i], conns[i+1:]...)
			// 		break
			// 	}
			// }
			broadcast("* disconnected " + conn.RemoteAddr().String() + "\n")
			conn.Close()
			break
		}
		fmt.Println("[" + time.Now().Format("15:04:05") + "] <" + conn.RemoteAddr().String() + ">: " + msg)

		// conn.Write([]byte("[" + time.Now().Format("15:04:05") + "] <" + conn.RemoteAddr().String() + ">: " + msg))
		broadcast("[" + time.Now().Format("15:04:05") + "] <" + conn.RemoteAddr().String() + ">: " + msg)
	}
}

// Server compartment
func Server(port string) {
	fmt.Println("works srv async")
	fmt.Println(port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		conns = append(conns, conn)

		for _, v := range conns {
			fmt.Println(v.RemoteAddr().String())
		}

		broadcast("* connected " + conn.RemoteAddr().String() + "\n")
		go handleConnection(conn)
	}

}
