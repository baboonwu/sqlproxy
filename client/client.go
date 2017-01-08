package main

import (
	"log"

	"github.com/siddontang/go-mysql/client"
	_ "github.com/siddontang/go-mysql/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Connect MySQL at 127.0.0.1:3306, with user root, an empty passowrd and database test
	conn, err := client.Connect("127.0.0.1:4000", "root", "", "")
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		conn.Close()
	}()

	conn.Ping()

	// Insert
	r, err := conn.Execute(`insert into customers (id, name, age) values (5, "abc", 14)`)
	if err != nil {
		log.Println(err)
	}
	if r != nil {
		// Get last insert id
		log.Println(r)
	}

	// Select
	result, err := conn.Execute(`select id, name, city from users where id > 0`)
	if err != nil {
		log.Println(err)
	}
	if result != nil {
		rowNum := result.RowNumber()
		for i := 0; i < rowNum; i++ {
			id, _ := result.GetStringByName(i, "id")
			name, _ := result.GetStringByName(i, "name")
			city, _ := result.GetStringByName(i, "city")
			log.Printf("row_%v:(%v %v %v) \n", i, id, name, city)
		}
	}

	// log.Println(result.Fields)
	// log.Println(result.Values)
	// log.Println(result.FieldNames)
	// log.Println(result.RowDatas)
	// log.Println(result.RowNumber())

}
