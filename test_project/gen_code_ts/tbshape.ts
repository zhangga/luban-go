export namespace demo {

    export class TbShape {
    private readonly _dataMap: Map<number, Shape>;
    private readonly _dataList: Shape[];
    
    constructor(dataList: Shape[]) {
        this._dataList = dataList;
        this._dataMap = new Map<number, Shape>();
        for (const v of this._dataList) {
            this._dataMap.set(v.Id, v);
        }
    }
    
    public get dataMap(): Map<number, Shape> {
        return this._dataMap;
    }
    
    public get dataList(): Shape[] {
        return this._dataList;
    }
    
    public get(key: number): Shape | undefined {
        return this._dataMap.get(key);
    }
}

}
