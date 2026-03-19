namespace demo
{

    public class TbItem
    {
        private readonly System.Collections.Generic.Dictionary<int, Item> _dataMap;
        private readonly System.Collections.Generic.List<Item> _dataList;
        
        public TbItem(System.Collections.Generic.List<Item> dataList)
        {
            _dataList = dataList;
            _dataMap = new System.Collections.Generic.Dictionary<int, Item>();
            foreach(var v in _dataList)
            {
                _dataMap.Add(v.Id, v);
            }
        }
        
        public System.Collections.Generic.Dictionary<int, Item> DataMap => _dataMap;
        public System.Collections.Generic.List<Item> DataList => _dataList;
        
        public Item GetOrDefault(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            return null;
        }
        
        public Item Get(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            throw new System.Collections.Generic.KeyNotFoundException();
        }
    }

}
