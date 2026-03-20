#pragma once

#include <vector>
#include <map>
#include <memory>


namespace demo {


class TbItem {
private:
    std::map<int32_t, std::shared_ptr<Item>> _dataMap;
    std::vector<std::shared_ptr<Item>> _dataList;

public:
    TbItem(const std::vector<std::shared_ptr<Item>>& dataList) {
        _dataList = dataList;
        for (auto& v : _dataList) {
            _dataMap[v->Id] = v;
        }
    }

    const std::map<int32_t, std::shared_ptr<Item>>& getDataMap() const {
        return _dataMap;
    }

    const std::vector<std::shared_ptr<Item>>& getDataList() const {
        return _dataList;
    }

    std::shared_ptr<Item> get(int32_t key) const {
        auto it = _dataMap.find(key);
        if (it != _dataMap.end()) {
            return it->second;
        }
        return nullptr;
    }
};


}
