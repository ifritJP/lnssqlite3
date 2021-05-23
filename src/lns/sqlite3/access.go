package sqlite3

import (
	"database/sql"
	//	"database/sql/driver"
	//"fmt"
	"log"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
	_ "github.com/mattn/go-sqlite3"
	//	"os"
)

type DB struct {
	tx    *sql.Tx
	sqldb *sql.DB
}

func callErrHanlde(_env *LnsEnv, errHandle LnsAny, stmt string, err error) {
	if !Lns_IsNil(errHandle) {
		errHandleFunc := errHandle.(Base_errHandleForm)
		errHandleFunc(_env,stmt, err.Error())
	} else {
		log.Fatal(err)
	}
}

func Open(_env *LnsEnv, path string, readonly bool, onMemoryFlag bool) (LnsAny, string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err.Error()
	}

	return &DB{nil, db}, "ok"
}

func (db *DB) Close(_env *LnsEnv) {
	db.sqldb.Close()
}

func (db *DB) Exec(_env *LnsEnv, stmt string, errHandle LnsAny) {
	var err error
	if db.tx == nil {
		_, err = db.sqldb.Exec(stmt)
	} else {
		var preStmt *sql.Stmt
		preStmt, err = db.tx.Prepare(stmt)
		if err == nil {
			defer preStmt.Close()
			preStmt.Exec()
		}
	}
	if err != nil {
		callErrHanlde(_env, errHandle, stmt, err)
	}
}

func (db *DB) Begin(_env *LnsEnv) {
	tx, err := db.sqldb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	db.tx = tx
}

func (db *DB) Commit(_env *LnsEnv) {
	if db.tx != nil {
		db.tx.Commit()
		db.tx = nil
	}
}

func (db *DB) query(_env *LnsEnv, query string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if db.tx != nil {
		rows, err = db.tx.Query(query)
	} else {
		rows, err = db.sqldb.Query(query)
	}
	return rows, err
}

func (db *DB) MapQuery(_env *LnsEnv, query string, callback LnsAny, errHandle LnsAny) bool {
	rows, err := db.query(_env, query)
	if err != nil {
		callErrHanlde(_env, errHandle, query, err)
		return false
	}
	defer rows.Close()
	columnNames, _ := rows.Columns()
	hasRow := false
	columns := make([]LnsAny, len(columnNames))
	columnsBuf := make([]LnsAny, len(columnNames))
	for index := 0; index < len(columnNames); index++ {
		var val LnsAny = nil
		columnsBuf[index] = &val
	}
	for rows.Next() {
		if !hasRow {
			hasRow = true
		}
		if err := rows.Scan(columnsBuf...); err != nil {
			callErrHanlde(_env, errHandle, query, err)
			break
		}
		for index := 0; index < len(columnNames); index++ {
			ifVal := *(columnsBuf[index].(*LnsAny))
			switch ifVal.(type) {
			case int64:
				columns[index] = LnsInt(ifVal.(int64))
			case []byte:
				columns[index] = string(ifVal.([]byte))
			default:
				columns[index] = ifVal
			}
		}

		if callback != nil {
			callbackFunc := callback.(Base_queryForm)
			if !callbackFunc(_env, NewLnsList(columns)) {
				break
			}
		}
	}
	err = rows.Err()
	if err != nil {
		callErrHanlde(_env, errHandle, query, err)
	}
	return hasRow
}

func (db *DB) MapQueryAsMap(
    _env *LnsEnv, query string, callback LnsAny, errHandle LnsAny) bool {
	rows, err := db.query(_env, query)
	if err != nil {
		callErrHanlde(_env, errHandle, query, err)
		return false
	}
	defer rows.Close()
	columnNames, _ := rows.Columns()
	hasRow := false
	columns := map[LnsAny]LnsAny{}
	columnsBuf := make([]LnsAny, len(columnNames))
	for index := 0; index < len(columnNames); index++ {
		var val LnsAny = nil
		columnsBuf[index] = &val
	}
	for rows.Next() {
		if !hasRow {
			hasRow = true
		}
		if err := rows.Scan(columnsBuf...); err != nil {
			callErrHanlde(_env, errHandle, query, err)
			break
		}
		for index := 0; index < len(columnNames); index++ {
			name := columnNames[index]
			ifVal := *(columnsBuf[index].(*LnsAny))
			switch ifVal.(type) {
			case int64:
				columns[name] = LnsInt(ifVal.(int64))
			case []byte:
				columns[name] = string(ifVal.([]byte))
			default:
				columns[name] = ifVal
			}
		}

		if callback != nil {
			callbackFunc := callback.(Base_queryMapForm)
			if !callbackFunc(_env, NewLnsMap(columns)) {
				break
			}
		}
	}
	err = rows.Err()
	if err != nil {
		callErrHanlde(_env, errHandle, query, err)
	}
	return hasRow
}
