// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_test bool
var test__mod__ string
// for 6
func test_convExp39(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func __main___anonymous_1010_(_env *LnsEnv, row *LnsList) bool {
    Lns_print([]LnsAny{Lns_forceCastInt(row.GetAt(1)) + 10, row.GetAt(2).(string) + "hoge"})
    return true
}
func __main___anonymous_1015_(_env *LnsEnv, row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1)})
    return false
}
func __main___anonymous_1020_(_env *LnsEnv, row *LnsMap) bool {
    Lns_print([]LnsAny{row.Get("id"), row.Get("name")})
    return true
}
// 4: decl @test.__main
func Test___main(_env *LnsEnv, args *LnsList) LnsInt {
    Lns_test_init( _env )
    var db base.Base_DB
    
    {
        _db := test_convExp39(Lns_2DDD(base.Base_Open(_env, "hoge.sqlite3", false, false)))
        if _db == nil{
            Lns_print([]LnsAny{"open error"})
            return 1
        } else {
            db = _db.(base.Base_DB)
        }
    }
    var stmt string
    stmt = "      create table foo (id integer not null primary key, name text);\n   delete from foo;\n"
    db.Exec(_env, stmt, nil)
    db.Begin(_env)
    {
        var _from89 LnsInt = 0
        var _to89 LnsInt = 10
        for _work89 := _from89; _work89 <= _to89; _work89++ {
            index := _work89
            var sql string
            sql = _env.LuaVM.String_format("insert into foo(id, name) values(%d, 'こんにちわ世界%03d')", []LnsAny{index, index})
            db.Exec(_env, sql, nil)
        }
    }
    db.Commit(_env)
    db.MapQuery(_env, "select id, name from foo", base.Base_queryForm(__main___anonymous_1010_), nil)
    db.MapQuery(_env, "select name from foo where id = 3", base.Base_queryForm(__main___anonymous_1015_), nil)
    db.Exec(_env, "delete from foo", nil)
    db.Exec(_env, "insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)
    db.MapQueryAsMap(_env, "select id, name from foo", base.Base_queryMapForm(__main___anonymous_1020_), nil)
    db.Close(_env)
    return 0
}

func Lns_test_init(_env *LnsEnv) {
    if init_test { return }
    init_test = true
    test__mod__ = "@test"
    Lns_InitMod()
    base.Lns_base_init(_env)
}
func init() {
    init_test = false
}
