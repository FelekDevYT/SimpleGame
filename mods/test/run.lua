local m = require("api")
local worldEvent = m.getWorldEvent()

if worldEvent.isWorldOpened() then
    m.addBlock(255,153,0)
end
