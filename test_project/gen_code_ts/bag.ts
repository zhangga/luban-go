export namespace demo {


    export class Bag {


    public readonly Id: number;


    public readonly Name: string;


    public readonly PriceList: number[];


    public readonly PropMap: Map<string, number>;


    constructor(Id: number, Name: string, PriceList: number[], PropMap: Map<string, number>) {


        this.Id = Id;

        this.Name = Name;

        this.PriceList = PriceList;

        this.PropMap = PropMap;

    }
}

}
