{{if .Namespace}}export namespace {{.Namespace}} {
{{end}}
{{if .Namespace}}    {{end}}export class {{.Name}} {
    private readonly _dataMap: Map<{{type .IndexType}}, {{type .ValueType}}>;
    private readonly _dataList: {{type .ValueType}}[];
    
    constructor(dataList: {{type .ValueType}}[]) {
        this._dataList = dataList;
        this._dataMap = new Map<{{type .IndexType}}, {{type .ValueType}}>();
        for (const v of this._dataList) {
            this._dataMap.set(v.{{.Index}}, v);
        }
    }
    
    public get dataMap(): Map<{{type .IndexType}}, {{type .ValueType}}> {
        return this._dataMap;
    }
    
    public get dataList(): {{type .ValueType}}[] {
        return this._dataList;
    }
    
    public get(key: {{type .IndexType}}): {{type .ValueType}} | undefined {
        return this._dataMap.get(key);
    }
}
{{if .Namespace}}
}
{{end}}