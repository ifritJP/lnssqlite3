// -*- coding: utf-8 -*-

package main

import (
	"fmt"
	"log"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
	lnssqlite3 "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
)

func initLnsRuntime() *LnsEnv {

	Lns_InitModOnce(LnsRuntimeOpt{})

	env := Lns_GetEnv()

	return env
}

func main() {

	env := initLnsRuntime()

	db_, err := lnssqlite3.Open(env, "hoge.sqlite3", false, false)
	log.Printf("%s", err)
	db := db_.(*lnssqlite3.DB)

	stmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	db.Exec(env, stmt, nil)

	db.Begin(env)
	for i := 0; i < 10; i++ {
		stmt := fmt.Sprintf(
			"insert into foo(id, name) values(%d, '%03d')", i, i)
		db.Exec(env, stmt, nil)
	}
	db.Commit(env)

	db.MapQuery(env, "select id, name from foo",
		lnssqlite3.Base_queryForm(func(env *LnsEnv, arg1 *LnsList) bool {
			fmt.Println(arg1.Items[0], arg1.Items[1])
			return true
		}), nil)

	db.MapQuery(
		env,
		"select name from foo where id = 3",
		lnssqlite3.Base_queryForm(func(env *LnsEnv, arg1 *LnsList) bool {
			fmt.Println(arg1.Items[0])
			return false
		}), nil)

	db.Exec(env, "delete from foo", nil)

	db.Exec(env, "insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)

	db.MapQuery(env, "select id, name from foo",
		lnssqlite3.Base_queryForm(func(env *LnsEnv, arg1 *LnsList) bool {
			fmt.Println(arg1.Items[0], arg1.Items[1])
			return true
		}), nil)

	db.Close(env)
}
