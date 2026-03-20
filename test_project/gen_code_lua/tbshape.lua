-- Table TbShape
local TbShape = {
    _dataMap = {},
    _dataList = {}
}

function TbShape.Load(dataList)
    TbShape._dataList = dataList
    for _, v in ipairs(dataList) do
        TbShape._dataMap[v.Id] = v
    end
end

function TbShape.Get(key)
    return TbShape._dataMap[key]
end

function TbShape.GetDataList()
    return TbShape._dataList
end

function TbShape.GetDataMap()
    return TbShape._dataMap
end

return TbShape