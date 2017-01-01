package main

import (
	"log"

	"github.com/siddontang/go-mysql/client"
)

func main() {

	// Connect MySQL at 127.0.0.1:3306, with user root, an empty passowrd and database test
	conn, err := client.Connect("127.0.0.1:4000", "root", "", "test")
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		conn.Close()
	}()

	conn.Ping()

	// Insert
	r, err := conn.Execute(`insert into table (id, name) values (1, "abc")`)
	if err != nil {
		conn.Close()
		log.Panic(err)
	}

	// Get last insert id
	println(r.InsertId)

	// Select
	r, _ = conn.Execute(`select id, name from table where id = 1`)

	// Handle resultset
	// v, _ := r.GetInt(0, 0)
	// v, _ = r.GetIntByName(0, "id")
}
