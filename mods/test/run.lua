local m = require("api")
local block = m.getBlockEvent()

m.set_block(1,1,0)

if block.isBlockSet() == true then
    print(block.isBlockSet())
    cords = block.getBlockCords()
    print(cords.x)
end
