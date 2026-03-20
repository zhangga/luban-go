#pragma once

#include <vector>
#include <map>
#include <memory>

{{if .Namespace}}
namespace {{.Namespace}} {
{{end}}

class {{.Name}} {
private:
    std::map<{{type .IndexType}}, std::shared_ptr<{{type .ValueType}}>> _dataMap;
    std::vector<std::shared_ptr<{{type .ValueType}}>> _dataList;

public:
    {{.Name}}(const std::vector<std::shared_ptr<{{type .ValueType}}>>& dataList) {
        _dataList = dataList;
        for (auto& v : _dataList) {
            _dataMap[v->{{.Index}}] = v;
        }
    }

    const std::map<{{type .IndexType}}, std::shared_ptr<{{type .ValueType}}>>& getDataMap() const {
        return _dataMap;
    }

    const std::vector<std::shared_ptr<{{type .ValueType}}>>& getDataList() const {
        return _dataList;
    }

    std::shared_ptr<{{type .ValueType}}> get({{type .IndexType}} key) const {
        auto it = _dataMap.find(key);
        if (it != _dataMap.end()) {
            return it->second;
        }
        return nullptr;
    }
};

{{if .Namespace}}
}
{{end}}