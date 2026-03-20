{{if .Namespace}}export namespace {{.Namespace}} {
{{end}}
{{if .Comment}}/**
 * {{.Comment}}
 */{{end}}
{{if .Namespace}}    {{end}}export {{if .IsDynamic}}abstract {{end}}class {{.Name}}{{if .Parent}} extends {{.Parent}}{{end}} {
{{range .ExportFields}}
{{if .Comment}}    /**
     * {{.Comment}}
     */{{end}}
    public readonly {{.Name}}: {{type .Type}};
{{end}}

    constructor({{- range $i, $f := .HierarchyFields}}{{if $i}}, {{end}}{{.Name}}: {{type .Type}}{{end -}}) {
{{if .Parent}}        super({{- range $i, $f := .ParentFields}}{{if $i}}, {{end}}{{.Name}}{{end -}});{{end}}
{{range .ExportFields}}
        this.{{.Name}} = {{.Name}};
{{end}}
    }
}
{{if .Namespace}}
}
{{end}}