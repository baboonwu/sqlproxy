package server

import (
	"database/sql"
	"log"
	"net"

	siddon "github.com/siddontang/go-mysql/server"
)

// StartProxyServer start tcp proxy server for Mysql.
func StartProxyServer(host string, db *sql.DB) {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	log.Println("start server successful", host)

	// begin to receive request
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go proxyHandle(conn, db)
	}
}

func proxyHandle(conn net.Conn, db *sql.DB) {
	// close connection before exit
	defer conn.Close()

	log.Println("recv client", conn.RemoteAddr().String())

	// Create a connection with user root and an empty passowrd
	// We only an empty handler to handle command too
	siddonconn, _ := siddon.NewConn(conn, "root", "", MysqlHandler{db: db})

	for {
		err := siddonconn.HandleCommand()
		if err != nil {
			log.Println(err)
			return
		}
	}
}
