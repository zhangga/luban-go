export namespace demo {

    export class TbItem {
    private readonly _dataMap: Map<number, Item>;
    private readonly _dataList: Item[];
    
    constructor(dataList: Item[]) {
        this._dataList = dataList;
        this._dataMap = new Map<number, Item>();
        for (const v of this._dataList) {
            this._dataMap.set(v.Id, v);
        }
    }
    
    public get dataMap(): Map<number, Item> {
        return this._dataMap;
    }
    
    public get dataList(): Item[] {
        return this._dataList;
    }
    
    public get(key: number): Item | undefined {
        return this._dataMap.get(key);
    }
}

}
