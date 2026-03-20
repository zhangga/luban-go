#pragma once

#include <vector>
#include <map>
#include <memory>


namespace demo {


class TbBag {
private:
    std::map<int32_t, std::shared_ptr<Bag>> _dataMap;
    std::vector<std::shared_ptr<Bag>> _dataList;

public:
    TbBag(const std::vector<std::shared_ptr<Bag>>& dataList) {
        _dataList = dataList;
        for (auto& v : _dataList) {
            _dataMap[v->Id] = v;
        }
    }

    const std::map<int32_t, std::shared_ptr<Bag>>& getDataMap() const {
        return _dataMap;
    }

    const std::vector<std::shared_ptr<Bag>>& getDataList() const {
        return _dataList;
    }

    std::shared_ptr<Bag> get(int32_t key) const {
        auto it = _dataMap.find(key);
        if (it != _dataMap.end()) {
            return it->second;
        }
        return nullptr;
    }
};


}
