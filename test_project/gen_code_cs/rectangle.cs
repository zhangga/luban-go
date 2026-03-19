namespace demo
{

    
    public class Rectangle : Shape
    {
        
        
        public readonly float Width;
        
        
        public readonly float Height;
        
        
        public Rectangle(int id, float width, float height) : base(id)
        {
            
            this.Width = width;
            
            this.Height = height;
            
        }
    }

}
