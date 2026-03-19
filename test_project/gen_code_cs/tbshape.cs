namespace demo
{

    public class TbShape
    {
        private readonly System.Collections.Generic.Dictionary<int, Shape> _dataMap;
        private readonly System.Collections.Generic.List<Shape> _dataList;
        
        public TbShape(System.Collections.Generic.List<Shape> dataList)
        {
            _dataList = dataList;
            _dataMap = new System.Collections.Generic.Dictionary<int, Shape>();
            foreach(var v in _dataList)
            {
                _dataMap.Add(v.Id, v);
            }
        }
        
        public System.Collections.Generic.Dictionary<int, Shape> DataMap => _dataMap;
        public System.Collections.Generic.List<Shape> DataList => _dataList;
        
        public Shape GetOrDefault(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            return null;
        }
        
        public Shape Get(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            throw new System.Collections.Generic.KeyNotFoundException();
        }
    }

}
