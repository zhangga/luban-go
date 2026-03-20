-- Table TbBag
local TbBag = {
    _dataMap = {},
    _dataList = {}
}

function TbBag.Load(dataList)
    TbBag._dataList = dataList
    for _, v in ipairs(dataList) do
        TbBag._dataMap[v.Id] = v
    end
end

function TbBag.Get(key)
    return TbBag._dataMap[key]
end

function TbBag.GetDataList()
    return TbBag._dataList
end

function TbBag.GetDataMap()
    return TbBag._dataMap
end

return TbBag