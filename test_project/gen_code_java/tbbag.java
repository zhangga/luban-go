package demo;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class TbBag {
    private final Map<int, Bag> _dataMap;
    private final List<Bag> _dataList;

    public TbBag(List<Bag> dataList) {
        this._dataList = dataList;
        this._dataMap = new HashMap<>();
        for (Bag v : _dataList) {
            this._dataMap.put(v.Id, v);
        }
    }

    public Map<int, Bag> getDataMap() {
        return _dataMap;
    }

    public List<Bag> getDataList() {
        return _dataList;
    }

    public Bag get(int key) {
        return _dataMap.get(key);
    }
}
