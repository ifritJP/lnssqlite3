module test

go 1.16

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20210705223321-4ff5d2c19c80
	github.com/ifritJP/lnssqlite3/src v0.0.0-20210415125250-f4bcfc7c361a
)

replace github.com/ifritJP/LuneScript/src => ../../LuneScript/src

replace github.com/ifritJP/lnssqlite3/src => ../src
