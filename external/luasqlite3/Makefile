LUA_COMMAND=lua5.3
LUA_INC=/usr/include/lua5.3

SO=so
INSTALL=install

LUASQLITE3_FILE=lsqlite3_fsl09y
LUASQLITE3_WORKDIR=.
LUASQLITE3_DIR=$(LUASQLITE3_WORKDIR)/$(LUASQLITE3_FILE)
#USE_BDB_SQLITE=y

LUASQLITE3_URL="http://lua.sqlite.org/index.cgi/zip/$(LUASQLITE3_FILE).zip?uuid=fsl_9y"

LUA_DIR_LIB="for path in string.gmatch( package.cpath, '[^;]+' ) do if path:find( \".*lua.*%?%.$(SO)\" ) then print( (path:gsub( \"/%?%.$(SO)\", '' ))); break end end"
LUA_LIB_DIR=$(shell echo $(LUA_DIR_LIB) | $(LUA_COMMAND) )

help:
	@echo make download
	@echo make build
	@echo make install


download:
	@echo check luasqlite3 ...
ifeq ($(wildcard $(LUASQLITE3_DIR).zip),)
	@echo install $(luasqlite3)
	mkdir -p $(LUASQLITE3_WORKDIR)
	curl --proxy "$(PROXY)" $(LUASQLITE3_URL) -o $(LUASQLITE3_DIR).zip
	(cd $(LUASQLITE3_WORKDIR); unzip -o $(shell basename $(LUASQLITE3_DIR).zip))
else
	@echo ready luasqlite3
endif

build: lsqlite3.$(SO)


ifeq ($(shell $(LUA_COMMAND) -e 'print( require( "lsqlite3" ) or "" )' 2> /dev/null),)
install: lsqlite3.$(SO)
	if [ -f lsqlite3.$(SO)  ]; then $(INSTALL) lsqlite3.$(SO) $(LUA_LIB_DIR); fi

lsqlite3.$(SO): $(LUASQLITE3_DIR)/lsqlite3.c $(LUASQLITE3_DIR)/sqlite3.c
ifeq ($(USE_BDB_SQLITE),y)
	gcc -std=c99 -fPIC -shared $(OPT_OP) -o lsqlite3.$(SO) \
		$(LUASQLITE3_DIR)/lsqlite3.c lsqlite3.dummy.c \
		-I$(LUA_INC) $(LUA_LIB) $(LUA_SO) -lpthread -ldb_sql-5.3
else
	gcc -std=c99 -fPIC -shared $(OPT_OP) -o lsqlite3.$(SO) \
		$(LUASQLITE3_DIR)/lsqlite3.c $(LUASQLITE3_DIR)/sqlite3.c \
		-I$(LUA_INC) $(LUA_LIB) $(LUA_SO) -lpthread
endif
else
install:
	@echo "lsqlite3 is ready for '$(LUA_COMMAND)'"
endif

