export namespace demo {


    export class Rectangle extends Shape {


    public readonly Width: number;


    public readonly Height: number;


    constructor(Id: number, Width: number, Height: number) {
        super(Id);

        this.Width = Width;

        this.Height = Height;

    }
}

}
