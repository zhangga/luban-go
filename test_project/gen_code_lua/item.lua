-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- Item 
local Item = {}




function Item.New(id, name, desc)
    local obj = {}
    
    
    obj.Id = id 
    
    obj.Name = name 
    
    obj.Desc = desc 
    
    return obj
end

return Item