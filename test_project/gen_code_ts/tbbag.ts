export namespace demo {

    export class TbBag {
    private readonly _dataMap: Map<number, Bag>;
    private readonly _dataList: Bag[];
    
    constructor(dataList: Bag[]) {
        this._dataList = dataList;
        this._dataMap = new Map<number, Bag>();
        for (const v of this._dataList) {
            this._dataMap.set(v.Id, v);
        }
    }
    
    public get dataMap(): Map<number, Bag> {
        return this._dataMap;
    }
    
    public get dataList(): Bag[] {
        return this._dataList;
    }
    
    public get(key: number): Bag | undefined {
        return this._dataMap.get(key);
    }
}

}
