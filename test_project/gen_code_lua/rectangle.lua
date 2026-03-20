-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- Rectangle 
local Rectangle = {}



-- Inherits from Shape


function Rectangle.New(id, width, height)
    local obj = {}
    
    -- inherit fields (assuming simple copy or metatable here, using simple assignment for generated struct)
    
    
    obj.Id = id 
    
    obj.Width = width 
    
    obj.Height = height 
    
    return obj
end

return Rectangle