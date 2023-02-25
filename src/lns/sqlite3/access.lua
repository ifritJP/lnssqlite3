--lns/sqlite3/access.lns
local _moduleObj = {}
local __mod__ = '@lns.@sqlite3.@access'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
end

function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
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

if not _lune8 then
   _lune8 = _lune
end




local function log( level, ... )

end

local NativeDB = {}


function NativeDB._setmeta( obj )
  setmetatable( obj, { __index = NativeDB  } )
end
function NativeDB._new( nrows, rows, exec, errmsg )
   local obj = {}
   NativeDB._setmeta( obj )
   if obj.__init then
      obj:__init( nrows, rows, exec, errmsg )
   end
   return obj
end
function NativeDB:__init( nrows, rows, exec, errmsg )

   self.nrows = nrows
   self.rows = rows
   self.exec = exec
   self.errmsg = errmsg
end


local lsqlite3 = require( "lsqlite3" )

local DBAccess = {}
_moduleObj.DBAccess = DBAccess
function DBAccess._new( db, readonly )
   local obj = {}
   DBAccess._setmeta( obj )
   if obj.__init then obj:__init( db, readonly ); end
   return obj
end
function DBAccess:__init(db, readonly) 
   self.db = db
   self.readonly = false
   self.beginFlag = false
end
function DBAccess._setmeta( obj )
  setmetatable( obj, { __index = DBAccess  } )
end


function DBAccess:errorExit( level, ... )

   print( "Sqlite ERROR:", self.db.errmsg( self.db ), ... )
   os.exit( 1 )
end


function DBAccess:Close(  )

   self.db:close(  )
end


local ActResult = {}
ActResult._name2Val = {}
function ActResult:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "ActResult.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function ActResult._from( val )
   return _lune._AlgeFrom( ActResult, val )
end

ActResult.Ng = { "Ng", {{ func=_lune._toStr, nilable=false, child={} }}}
ActResult._name2Val["Ng"] = ActResult.Ng
ActResult.Ok = { "Ok", {{ func=_lune._toList, nilable=false, child={ { func = _lune._toStem, nilable = false, child = {} } } }}}
ActResult._name2Val["Ok"] = ActResult.Ok




local function createActForm(  )

   local loaded, message = _lune.loadstring52( [==[
return function( db, form, ... )
  local result
  local success, message = pcall( function(...)
    result = { form( db, ... ) }
  end, ... )
  if success then
    return success, result
  end
  return success, message
end
]==] )
   if loaded ~= nil then
      local funcObj = _lune.unwrap( loaded(  ))
      return funcObj
   else
      error( message or "load message" )
   end
   
end

local actForm = createActForm(  )

function DBAccess:act( dbForm, ... )

   local success, result = actForm( self.db, dbForm, ... )
   if success then
      return _lune.newAlge( ActResult.Ok, {result})
   else
    
      return _lune.newAlge( ActResult.Ng, {result})
   end
   
end


function DBAccess:Exec( stmt, errHandle )

   local errmsg
   
   do
      local _matchExp = self:act( self.db.exec, stmt )
      if _matchExp[1] == ActResult.Ok[1] then
         local list = _matchExp[2][1]
      
         if list[1] ~= lsqlite3.OK then
            do
               local _matchExp = self:act( self.db.errmsg )
               if _matchExp[1] == ActResult.Ok[1] then
                  local list2 = _matchExp[2][1]
               
                  errmsg = list2[1]
               elseif _matchExp[1] == ActResult.Ng[1] then
                  local msg2 = _matchExp[2][1]
               
                  errmsg = msg2
               end
            end
            
         else
          
            errmsg = nil
         end
         
      elseif _matchExp[1] == ActResult.Ng[1] then
         local msg = _matchExp[2][1]
      
         errmsg = msg
      end
   end
   
   if errmsg ~= nil then
      if errHandle ~= nil then
         errHandle( stmt, errmsg )
      else
         self:errorExit( 3, stmt )
      end
      
   end
   
end


function DBAccess:Begin(  )

   if self.readonly then
      print( 1, "db mode is read only" )
      os.exit( 1 )
   end
   
   self:Exec( "PRAGMA journal_mode = MEMORY" )
   self:Exec( "PRAGMA synchronous = OFF" )
   self:Exec( "BEGIN IMMEDIATE" )
   
   self.beginFlag = true
end


function DBAccess:Commit(  )

   if self.readonly then
      return 
   end
   
   if not self.beginFlag then
      return 
   end
   
   self.beginFlag = false
   
   self:Exec( "COMMIT", function ( stmt, msg )
   
      if not msg:find( "no transaction is active" ) then
         self:errorExit( 5, msg )
      end
      
      
   end )
end


function DBAccess:MapQuery( query, callback, errHandle )

   local hit = false
   
   while true do
      
      do
         local _matchExp = self:act( self.db.rows, query )
         if _matchExp[1] == ActResult.Ng[1] then
            local msg = _matchExp[2][1]
         
            if not msg:find( " is locked", 1, true ) then
               print( 1, msg )
               
               break
            end
            
            self:errorExit( 3, msg, query )
         elseif _matchExp[1] == ActResult.Ok[1] then
            local list = _matchExp[2][1]
         
            local loopFunc, param, prev = table.unpack( list )
            if loopFunc ~= nil and callback ~= nil then
               while true do
                  local workItem = (loopFunc )( param, prev )
                  if  nil == workItem then
                     local _workItem = workItem
                  
                     break
                  end
                  
                  prev = workItem
                  if not hit then
                     hit = true
                  end
                  
                  local item = {}
                  for __index, val in pairs( workItem ) do
                     table.insert( item, val )
                  end
                  
                  if not callback( item ) then
                     break
                  end
                  
               end
               
            end
            
            break
         end
      end
      
   end
   
   return hit
end


function DBAccess:MapQueryAsMap( query, callback, errHandle )

   local hit = false
   
   while true do
      
      do
         local _matchExp = self:act( self.db.nrows, query )
         if _matchExp[1] == ActResult.Ng[1] then
            local msg = _matchExp[2][1]
         
            if not msg:find( " is locked", 1, true ) then
               print( 1, msg )
               
               break
            end
            
            self:errorExit( 3, msg, query )
         elseif _matchExp[1] == ActResult.Ok[1] then
            local list = _matchExp[2][1]
         
            local loopFunc, param, prev = table.unpack( list )
            if loopFunc ~= nil and callback ~= nil then
               while true do
                  local item = (loopFunc )( param, prev )
                  if  nil == item then
                     local _item = item
                  
                     break
                  end
                  
                  prev = item
                  if not hit then
                     hit = true
                  end
                  
                  if not callback( item ) then
                     break
                  end
                  
               end
               
            end
            
            break
         end
      end
      
   end
   
   return hit
end


local function Open( path, readonly, onMemoryFlag )

   local flag
   
   if readonly then
      flag = lsqlite3.OPEN_READONLY
   else
    
      flag = lsqlite3.OPEN_READWRITE + lsqlite3.OPEN_CREATE
   end
   
   log( 3, "DBAccess:open", flag )
   
   local db
   
   do
      local workDB
      
      if onMemoryFlag then
         workDB = lsqlite3.open_memory(  )
      else
       
         workDB = lsqlite3.open( path, flag )
      end
      
      if workDB ~= nil then
         db = workDB
      else
         log( 1, "open error." )
         return nil, ""
      end
      
   end
   
   
   return DBAccess._new(db, readonly), ""
end
_moduleObj.Open = Open

return _moduleObj
