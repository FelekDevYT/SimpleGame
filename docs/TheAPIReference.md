# The API reference of FelsAPI

## What is is, FelsAPI?

FelsAPI - API for you can mod the felsCraft<br>
felsAPI use Lua for working with game<br>
Lua version is Lua 5.2<br>
[Lua 5.2 docs](https://www.lua.org/manual/5.2/manual.html#6.9)<br>


## The API reference


### Written up the structure of modification

Create in main folder(main folder should have game.exe file) directory with name "mods"<br>
And next you should create folder with name is you mod name<br>
Create in you folder file with name "run.lua" and open it<br>

## Create the first mod

### The importing module

```Lua
local api = require("api")--load the API for working with events and blocks
```

### Events

### Importing the EventHandler

```Lua
local BlockEventHandler = api.getBlockEvent()
```

### The events

#### isBlockSet() event

```Lua
if BlockEventHandler.isBlockSet() == true then
    print("Event is called!")
end 
```

#### Getting Event args

```Lua
if BlockEventHandler.isBlockSet() == true then
    cords = BlockEventHandler.getBlockCords()--getting cords of event
    print("[ EVENT OUTPUT ] "+cords.x+", "+cords.y+". "+BlockEventHandler.getBlockType)
end 
```

### Standard operations

#### Set block in an pos

```Lua
api.set_block(1,1,4)--[1] - x,[2] - y,[3] - block type
```

#### getting block type from an pos

```Lua
print(api.get_block(1,1))--[1] - x,[2] - y
```
