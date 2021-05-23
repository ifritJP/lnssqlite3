// This code is transcompiled by LuneScript.
package sqlite3
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_base bool
var base__mod__ string
type Base_errHandleForm func (_env *LnsEnv, arg1 string,arg2 string)
type Base_queryForm func (_env *LnsEnv, arg1 *LnsList) bool
type Base_queryMapForm func (_env *LnsEnv, arg1 *LnsMap) bool
// 20: decl @lns.@sqlite3.@base.Open
func Base_Open(_env *LnsEnv, path string,readonly bool,onMemoryFlag bool)(LnsAny, string) {
    return Open(_env, path, readonly, onMemoryFlag)
}

type Base_DB interface {
        Begin(_env *LnsEnv)
        Close(_env *LnsEnv)
        Commit(_env *LnsEnv)
        Exec(_env *LnsEnv, arg1 string, arg2 LnsAny)
        MapQuery(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) bool
        MapQueryAsMap(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) bool
}
func Lns_cast2Base_DB( obj LnsAny ) LnsAny {
    if _, ok := obj.(Base_DB); ok { 
        return obj
    }
    return nil
}


func Lns_base_init(_env *LnsEnv) {
    if init_base { return }
    init_base = true
    base__mod__ = "@lns.@sqlite3.@base"
    Lns_InitMod()
}
func init() {
    init_base = false
}
