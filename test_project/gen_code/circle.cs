namespace demo
{

    
    public class Circle : Shape
    {
        
        
        public readonly float Radius;
        
        
        public Circle(int id, float radius) : base(id)
        {
            
            this.Radius = radius;
            
        }
    }

}
