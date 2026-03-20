package demo;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class TbItem {
    private final Map<int, Item> _dataMap;
    private final List<Item> _dataList;

    public TbItem(List<Item> dataList) {
        this._dataList = dataList;
        this._dataMap = new HashMap<>();
        for (Item v : _dataList) {
            this._dataMap.put(v.Id, v);
        }
    }

    public Map<int, Item> getDataMap() {
        return _dataMap;
    }

    public List<Item> getDataList() {
        return _dataList;
    }

    public Item get(int key) {
        return _dataMap.get(key);
    }
}
