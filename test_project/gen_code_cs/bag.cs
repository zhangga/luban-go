namespace demo
{

    
    public class Bag
    {
        
        
        public readonly int Id;
        
        
        public readonly string Name;
        
        
        public readonly System.Collections.Generic.List<int> PriceList;
        
        
        public readonly System.Collections.Generic.Dictionary<string, int> PropMap;
        
        
        public Bag(int id, string name, System.Collections.Generic.List<int> pricelist, System.Collections.Generic.Dictionary<string, int> propmap)
        {
            
            this.Id = id;
            
            this.Name = name;
            
            this.PriceList = pricelist;
            
            this.PropMap = propmap;
            
        }
    }

}
