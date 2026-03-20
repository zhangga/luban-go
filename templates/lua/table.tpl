-- Table {{.Name}}
local {{.Name}} = {
    _dataMap = {},
    _dataList = {}
}

function {{.Name}}.Load(dataList)
    {{.Name}}._dataList = dataList
    for _, v in ipairs(dataList) do
        {{.Name}}._dataMap[v.{{.Index}}] = v
    end
end

function {{.Name}}.Get(key)
    return {{.Name}}._dataMap[key]
end

function {{.Name}}.GetDataList()
    return {{.Name}}._dataList
end

function {{.Name}}.GetDataMap()
    return {{.Name}}._dataMap
end

return {{.Name}}