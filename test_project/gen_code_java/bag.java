package demo;


public class Bag {
    
    
    public final int Id;
    
    
    public final String Name;
    
    
    public final java.util.List<Integer> PriceList;
    
    
    public final java.util.Map<String, Integer> PropMap;
    

    public Bag(int id, String name, java.util.List<Integer> pricelist, java.util.Map<String, Integer> propmap) {
        
        
        this.Id = id;
        
        this.Name = name;
        
        this.PriceList = pricelist;
        
        this.PropMap = propmap;
        
    }
}
