-- Table TbItem
local TbItem = {
    _dataMap = {},
    _dataList = {}
}

function TbItem.Load(dataList)
    TbItem._dataList = dataList
    for _, v in ipairs(dataList) do
        TbItem._dataMap[v.Id] = v
    end
end

function TbItem.Get(key)
    return TbItem._dataMap[key]
end

function TbItem.GetDataList()
    return TbItem._dataList
end

function TbItem.GetDataMap()
    return TbItem._dataMap
end

return TbItem