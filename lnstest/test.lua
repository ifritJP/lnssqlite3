--test.lns
local _moduleObj = {}
local __mod__ = '@test'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune3 then
   _lune3 = _lune
end
local base = _lune.loadModule( 'go/github:com.ifritJP.lnssqlite3.base' )

local function __main( args )

   
   local db = base.Open( "hoge.sqlite3", false, false )
   if  nil == db then
      local _db = db
   
      print( "open error" )
      return 1
   end
   
   
   local stmt = [==[
      create table foo (id integer not null primary key, name text);
   delete from foo;
]==]
   db:Exec( stmt, nil )
   
   db:Begin(  )
   for index = 0, 10 do
      local sql = string.format( "insert into foo(id, name) values(%d, 'こんにちわ世界%03d')", index, index)
      db:Exec( sql, nil )
   end
   
   db:Commit(  )
   
   db:MapQuery( "select id, name from foo", function ( row )
   
      print( math.floor(row[1]) + 10, row[2] .. "hoge" )
      return true
   end )
   
   db:MapQuery( "select name from foo where id = 3", function ( row )
   
      print( row[1] )
      return false
   end )
   
   db:Exec( "delete from foo", nil )
   
   db:Exec( "insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil )
   
   db:MapQuery( "select id, name from foo", function ( row )
   
      print( row[1], row[2] )
      return true
   end )
   
   db:Close(  )
   
   return 0
end
_moduleObj.__main = __main

return _moduleObj
