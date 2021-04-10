package sqlite3

import (
	"database/sql"
    //	"database/sql/driver"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
    . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
	"log"
    //	"os"
)

type DB struct {
    tx *sql.Tx
    sqldb *sql.DB
}


func Open( path string, readonly bool, onMemoryFlag bool ) (LnsAny, string) {
	db, err := sql.Open("sqlite3", path )
	if err != nil {
		return nil, err.Error()
	}
    
    return &DB{ nil, db }, "ok"
}

func (db *DB) Close() {
    db.sqldb.Close()
}

func (db *DB) Exec( stmt string, errHandle LnsAny ) {
    var err error
    if db.tx == nil {
        _, err = db.sqldb.Exec(stmt)
    } else {
        var preStmt *sql.Stmt
        preStmt, err = db.tx.Prepare(stmt)
        defer preStmt.Close()
        preStmt.Exec()
    }
	if err != nil {
        if errHandle != nil {
            errHandleFunc := errHandle.(func( stmt string, msg string ))
            errHandleFunc( stmt, err.Error() )
        } else {
            log.Fatal( err )
        }
	}
}

func (db *DB) Begin() {
	tx, err := db.sqldb.Begin()
	if err != nil {
		log.Fatal(err)
	}
    db.tx = tx
}

func (db *DB) Commit() {
    if db.tx != nil {
        db.tx.Commit()
        db.tx = nil
    }
}


func (db *DB) MapQuery( query string, callback LnsAny ) bool {
	rows, err := db.sqldb.Query( query )
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
    columnNames, _ := rows.Columns()
    hasRow := false
    columns := make( []LnsAny, len( columnNames ) )
    columnsBuf := make( []LnsAny, len( columnNames ) )
    for index := 0; index < len( columnNames ); index++ {
        var val LnsAny = nil
        columnsBuf[ index ] = &val
    }
	for rows.Next() {
        if !hasRow {
            hasRow = true
        }
        if err := rows.Scan( columnsBuf... ); err != nil {
            log.Fatal( err )
        }
        for index := 0; index < len( columnNames ); index++ {
            ifVal := *(columnsBuf[index].(*LnsAny))
            switch ifVal.(type) {
            case int64:
                columns[ index ] = LnsInt( ifVal.(int64) )
            case []byte:
                columns[ index ] = string( ifVal.([]byte ) )
            default:
                columns[ index ] = ifVal
            }
        }

        if callback != nil {
            callbackFunc := callback.(Base_queryForm)
            if !callbackFunc( NewLnsList( columns ) ) {
                break
            }
        }
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
    return hasRow
}
