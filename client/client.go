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

	// Delete
	r, err := conn.Execute(`delete from customers where id=4`)
	if err != nil {
		log.Println(err)
	}
	if r != nil {
		log.Println(r)
	}

	// Insert
	r, err = conn.Execute(`insert into customers (id, name, age) values (4, "abc", 14)`)
	if err != nil {
		log.Println(err)
	}
	if r != nil {
		log.Println(r)
	}

	r, err = conn.Execute(`update customers set age=27 where id=4`)
	if err != nil {
		log.Println(err)
	}
	if r != nil {
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
