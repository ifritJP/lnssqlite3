--lns/sqlite3/base.lns
local _moduleObj = {}
local __mod__ = '@lns.@sqlite3.@base'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
if not _lune3 then
   _lune3 = _lune
end




local DB = {}
_moduleObj.DB = DB
function DB.setmeta( obj )
  setmetatable( obj, { __index = DB  } )
end
function DB.new(  )
   local obj = {}
   DB.setmeta( obj )
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
