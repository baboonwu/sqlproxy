package server

import (
	"fmt"
	"log"

	. "github.com/siddontang/go-mysql/mysql"
)

type MysqlHandler struct {
}

func (h MysqlHandler) UseDB(dbName string) error {
	log.Printf("use %s \n", dbName)
	return nil
}

func (h MysqlHandler) HandleQuery(query string) (*Result, error) {
	log.Printf("HandleQuery: %s \n", query)
	return &Result{0, 1, 0, nil}, nil
}

func (h MysqlHandler) HandleFieldList(table string, fieldWildcard string) ([]*Field, error) {
	return nil, fmt.Errorf("HandleFieldList: not supported now")
}

func (h MysqlHandler) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	return 0, 0, nil, fmt.Errorf("HandleStmtPrepare: not supported now")
}

func (h MysqlHandler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*Result, error) {
	return nil, fmt.Errorf("HandleStmtExecute: not supported now")
}

func (h MysqlHandler) HandleStmtClose(context interface{}) error {
	log.Printf("HandleStmtClose\n")
	return nil
}
