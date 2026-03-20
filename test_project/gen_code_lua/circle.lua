-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- Circle 
local Circle = {}



-- Inherits from Shape


function Circle.New(id, radius)
    local obj = {}
    
    -- inherit fields (assuming simple copy or metatable here, using simple assignment for generated struct)
    
    
    obj.Id = id 
    
    obj.Radius = radius 
    
    return obj
end

return Circle