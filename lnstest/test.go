// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_test bool
var test__mod__ string
// for 5
func test_convExp39(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func __main___anonymous_1014_(row *LnsList) bool {
    Lns_print([]LnsAny{Lns_forceCastInt(row.GetAt(1)) + 10, row.GetAt(2).(string) + "hoge"})
    return true
}
func __main___anonymous_1021_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1)})
    return false
}
func __main___anonymous_1028_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1), row.GetAt(2)})
    return true
}
// 3: decl @test.__main
func Test___main(args *LnsList) LnsInt {
    Lns_test_init()
    var db base.Base_DB
    
    {
        _db := test_convExp39(Lns_2DDD(base.Base_Open("hoge.sqlite3", false, false)))
        if _db == nil{
            Lns_print([]LnsAny{"open error"})
            return 1
        } else {
            db = _db.(base.Base_DB)
        }
    }
    var stmt string
    stmt = "      create table foo (id integer not null primary key, name text);\n   delete from foo;\n"
    db.Exec(stmt, nil)
    db.Begin()
    {
        var _from89 LnsInt = 0
        var _to89 LnsInt = 10
        for _work89 := _from89; _work89 <= _to89; _work89++ {
            index := _work89
            var sql string
            sql = Lns_getVM().String_format("insert into foo(id, name) values(%d, 'こんにちわ世界%03d')", []LnsAny{index, index})
            db.Exec(sql, nil)
        }
    }
    db.Commit()
    db.MapQuery("select id, name from foo", base.Base_queryForm(__main___anonymous_1014_))
    db.MapQuery("select name from foo where id = 3", base.Base_queryForm(__main___anonymous_1021_))
    db.Exec("delete from foo", nil)
    db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)
    db.MapQuery("select id, name from foo", base.Base_queryForm(__main___anonymous_1028_))
    db.Close()
    return 0
}

func Lns_test_init() {
    if init_test { return }
    init_test = true
    test__mod__ = "@test"
    Lns_InitMod()
    base.Lns_base_init()
}
func init() {
    init_test = false
}
