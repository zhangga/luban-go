namespace demo
{

    public class TbBag
    {
        private readonly System.Collections.Generic.Dictionary<int, Bag> _dataMap;
        private readonly System.Collections.Generic.List<Bag> _dataList;
        
        public TbBag(System.Collections.Generic.List<Bag> dataList)
        {
            _dataList = dataList;
            _dataMap = new System.Collections.Generic.Dictionary<int, Bag>();
            foreach(var v in _dataList)
            {
                _dataMap.Add(v.Id, v);
            }
        }
        
        public System.Collections.Generic.Dictionary<int, Bag> DataMap => _dataMap;
        public System.Collections.Generic.List<Bag> DataList => _dataList;
        
        public Bag GetOrDefault(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            return null;
        }
        
        public Bag Get(int key)
        {
            if (_dataMap.TryGetValue(key, out var v))
            {
                return v;
            }
            throw new System.Collections.Generic.KeyNotFoundException();
        }
    }

}
