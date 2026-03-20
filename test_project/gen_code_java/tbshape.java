package demo;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class TbShape {
    private final Map<int, Shape> _dataMap;
    private final List<Shape> _dataList;

    public TbShape(List<Shape> dataList) {
        this._dataList = dataList;
        this._dataMap = new HashMap<>();
        for (Shape v : _dataList) {
            this._dataMap.put(v.Id, v);
        }
    }

    public Map<int, Shape> getDataMap() {
        return _dataMap;
    }

    public List<Shape> getDataList() {
        return _dataList;
    }

    public Shape get(int key) {
        return _dataMap.get(key);
    }
}
