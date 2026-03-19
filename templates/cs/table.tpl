{{if .Namespace}}namespace {{.Namespace}}
{
{{end}}
    public class {{.Name}}
    {
        private readonly System.Collections.Generic.Dictionary<{{type .IndexType}}, {{type .ValueType}}> _dataMap;
        private readonly System.Collections.Generic.List<{{type .ValueType}}> _dataList;
        
        public {{.Name}}(System.Collections.Generic.List<{{type .ValueType}}> dataList)
        {
            _dataList = dataList;
            _dataMap = new System.Collections.Generic.Dictionary<{{type .IndexType}}, {{type .ValueType}}>();
            foreach(var v in _dataList)
            {
                _dataMap.Add(v.{{title .Index}}, v);
            }
        }
        
        public System.Collections.Generic.Dictionary<{{type .IndexType}}, {{type .ValueType}}> DataMap => _dataMap;
        public System.Collections.Generic.List<{{type .ValueType}}> DataList => _dataList;
        
        public {{type .ValueType}} GetOrDefault({{type .IndexType}} key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            return null;
        }
        
        public {{type .ValueType}} Get({{type .IndexType}} key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            throw new System.Collections.Generic.KeyNotFoundException();
        }
    }
{{if .Namespace}}
}
{{end}}