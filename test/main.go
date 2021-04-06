package main

import (
	"log"
	"fmt"
    "github.com/ifritJP/lnssqlite3"
    . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

func main() {
    db_, err := lnssqlite3.Open( "hoge.sqlite3", false, false )
    log.Printf( "%s", err )
    db := db_.(*lnssqlite3.DB)


    stmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
    db.Exec( stmt, nil )

    db.Begin()
	for i := 0; i < 10; i++ {
        stmt := fmt.Sprintf(
            "insert into foo(id, name) values(%d, 'こんにちわ世界%03d')", i, i)
		db.Exec( stmt, nil )
	}
    db.Commit()
    
    


    db.MapQuery( "select id, name from foo",
        func ( row []LnsAny ) bool {
            fmt.Println( row[0], row[1] )
            return true
        })

    db.MapQuery(
        "select name from foo where id = 3",
        func ( row []LnsAny ) bool {
            fmt.Println( row[0] )
            return false
        } )
    

	db.Exec("delete from foo", nil)

	db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)
    

    db.MapQuery( "select id, name from foo",
        func ( row []LnsAny ) bool {
            fmt.Println( row[0], row[1])
            return true
        } );
    

    
    db.Close()
}
