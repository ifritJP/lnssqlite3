module test

go 1.16

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20230225113348-72886d7e513d
	github.com/ifritJP/lnssqlite3/src v0.0.0-20220603234035-58c55d39410f
)

replace github.com/ifritJP/LuneScript/src => ../../LuneScript/src

replace github.com/ifritJP/lnssqlite3/src => ../src
