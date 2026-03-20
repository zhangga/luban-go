export namespace demo {


    export class Circle extends Shape {


    public readonly Radius: number;


    constructor(Id: number, Radius: number) {
        super(Id);

        this.Radius = Radius;

    }
}

}
