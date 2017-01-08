package main

import (
	"flag"
	"log"

	"database/sql"

	"github.com/baboonwu/sqlproxy/server"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	hostport := flag.String("host:port", "127.0.0.1:4000", "host:port")
	flag.Parse()

	// start
	db, err := sql.Open("mysql", "root:baboon@tcp(127.0.0.1:3306)/baboon?charset=utf8")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	server.StartProxyServer(*hostport, db)
}
