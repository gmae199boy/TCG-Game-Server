package main

import (
	"log"
	"net"

	dbserver "github.com/gmae199boy/tcgServer/DBServer"
	connectServer "github.com/gmae199boy/tcgServer/connect"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	if nil != err {
		log.Println(err)
	}
	defer l.Close()

	var c connectServer.Connection
	var b connectServer.BattleLobby
	var db dbserver.DB
	c.InitConn()
	b.InitBattleLobby()
	db.ConnectDB()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			continue
		}
		go c.ServerLobby(conn, b)
	}
}
