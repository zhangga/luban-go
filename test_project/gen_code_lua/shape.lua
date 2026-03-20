-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- Shape 
local Shape = {}


-- This is an abstract/dynamic class.



function Shape.New(id)
    local obj = {}
    
    
    obj.Id = id 
    
    return obj
end

return Shape