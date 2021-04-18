// This code is transcompiled by LuneScript.
package sqlite3
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_base bool
var base__mod__ string
type Base_errHandleForm func (arg1 string,arg2 string)
type Base_queryForm func (arg1 *LnsList) bool
type Base_queryMapForm func (arg1 *LnsMap) bool
// 18: decl @lns.@sqlite3.@base.Open
func Base_Open(path string,readonly bool,onMemoryFlag bool)(LnsAny, string) {
    return Open(path, readonly, onMemoryFlag)
}

type Base_DB interface {
        Begin()
        Close()
        Commit()
        Exec(arg1 string, arg2 LnsAny)
        MapQuery(arg1 string, arg2 LnsAny) bool
        MapQueryAsMap(arg1 string, arg2 LnsAny) bool
}
func Lns_cast2Base_DB( obj LnsAny ) LnsAny {
    if _, ok := obj.(Base_DB); ok { 
        return obj
    }
    return nil
}


func Lns_base_init() {
    if init_base { return }
    init_base = true
    base__mod__ = "@lns.@sqlite3.@base"
    Lns_InitMod()
}
func init() {
    init_base = false
}
