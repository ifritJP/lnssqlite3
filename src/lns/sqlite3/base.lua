--lns/sqlite3/base.lns
local _moduleObj = {}
local __mod__ = '@lns.@sqlite3.@base'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
if not _lune6 then
   _lune6 = _lune
end






local DB = {}
_moduleObj.DB = DB
function DB._setmeta( obj )
  setmetatable( obj, { __index = DB  } )
end
function DB._new(  )
   local obj = {}
   DB._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function DB:__init(  )

end


local Access = require( "lns.sqlite3.access" )

local function Open( path, readonly, onMemoryFlag )

   return Access.Open( path, readonly, onMemoryFlag )
end
_moduleObj.Open = Open

return _moduleObj
