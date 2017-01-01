package main

import (
	"flag"

	"github.com/baboonwu/sqlproxy/server"
)

func main() {
	hostport := flag.String("host:port", "127.0.0.1:4000", "host:port")
	flag.Parse()

	server.StartProxyServer(*hostport)
}
