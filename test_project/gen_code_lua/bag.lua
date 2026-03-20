-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- Bag 
local Bag = {}




function Bag.New(id, name, pricelist, propmap)
    local obj = {}
    
    
    obj.Id = id 
    
    obj.Name = name 
    
    obj.PriceList = pricelist 
    
    obj.PropMap = propmap 
    
    return obj
end

return Bag