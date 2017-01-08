package server

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/siddontang/go-mysql/mysql"
	"github.com/xwb1989/sqlparser"
)

type MysqlHandler struct {
	db *sql.DB
}

func (h MysqlHandler) UseDB(dbName string) error {
	log.Printf("use %s \n", dbName)
	return nil
}

// HandleQuery is response to process select/insert/update/delete statement
func (h MysqlHandler) HandleQuery(queryStr string) (*Result, error) {
	// 1. parse
	statement, err := sqlparser.Parse(queryStr)
	if err != nil {
		log.Printf("HandleQuery: %s, err=%v \n", queryStr, err)
		return nil, err
	}

	var ok bool
	var query *sqlparser.Select
	// var insert *sqlparser.Insert
	var result *Result

	// 2. convert type to Select/Insert/Upadte/Delete
	switch statement.(type) {
	case *sqlparser.Select:
		query, ok = statement.(*sqlparser.Select)
		if !ok {
			return nil, fmt.Errorf("convert to select sql failed. sql=%s \n", query)
		}
		result, err = h.handleSelect(query)

	// TODO: open comment for insert
	// case *sqlparser.Insert:
	// 	insert, ok = statement.(*sqlparser.Insert)
	// 	if !ok {
	// 		return nil, fmt.Errorf("convert to insert sql failed. sql=%s \n", insert)
	// 	}
	// 	// log.Println(sqlparser.String(insert.Table))
	// 	// result, err = handleSelect(query)

	default:

	}

	return result, err
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
