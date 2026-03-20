{{if .Namespace}}package {{.Namespace}};{{end}}

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class {{.Name}} {
    private final Map<{{type .IndexType}}, {{type .ValueType}}> _dataMap;
    private final List<{{type .ValueType}}> _dataList;

    public {{.Name}}(List<{{type .ValueType}}> dataList) {
        this._dataList = dataList;
        this._dataMap = new HashMap<>();
        for ({{type .ValueType}} v : _dataList) {
            this._dataMap.put(v.{{.Index}}, v);
        }
    }

    public Map<{{type .IndexType}}, {{type .ValueType}}> getDataMap() {
        return _dataMap;
    }

    public List<{{type .ValueType}}> getDataList() {
        return _dataList;
    }

    public {{type .ValueType}} get({{type .IndexType}} key) {
        return _dataMap.get(key);
    }
}
