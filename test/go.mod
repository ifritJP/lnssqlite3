module github.com/ifritJP/lnssqlite3/test

go 1.20

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20230225113348-72886d7e513d
	github.com/ifritJP/lnssqlite3/src v0.0.0-20220603234035-58c55d39410f
)

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/yuin/gopher-lua v0.0.0-20220413183635-c841877397d8 // indirect
)

replace github.com/ifritJP/lnssqlite3/src/lns/sqlite3 => ../src
