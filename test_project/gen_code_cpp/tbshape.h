#pragma once

#include <vector>
#include <map>
#include <memory>


namespace demo {


class TbShape {
private:
    std::map<int32_t, std::shared_ptr<Shape>> _dataMap;
    std::vector<std::shared_ptr<Shape>> _dataList;

public:
    TbShape(const std::vector<std::shared_ptr<Shape>>& dataList) {
        _dataList = dataList;
        for (auto& v : _dataList) {
            _dataMap[v->Id] = v;
        }
    }

    const std::map<int32_t, std::shared_ptr<Shape>>& getDataMap() const {
        return _dataMap;
    }

    const std::vector<std::shared_ptr<Shape>>& getDataList() const {
        return _dataList;
    }

    std::shared_ptr<Shape> get(int32_t key) const {
        auto it = _dataMap.find(key);
        if (it != _dataMap.end()) {
            return it->second;
        }
        return nullptr;
    }
};


}
