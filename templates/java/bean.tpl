{{if .Namespace}}package {{.Namespace}};{{end}}

{{if .Comment}}/**
 * {{.Comment}}
 */{{end}}
public {{if .IsDynamic}}abstract {{end}}class {{.Name}}{{if .Parent}} extends {{.Parent}}{{end}} {
    {{range .ExportFields}}
    {{if .Comment}}/**
     * {{.Comment}}
     */{{end}}
    public final {{type .Type}} {{.Name}};
    {{end}}

    public {{.Name}}({{- range $i, $f := .HierarchyFields}}{{if $i}}, {{end}}{{type $f.Type}} {{lower $f.Name}}{{end -}}) {
        {{if .Parent}}super({{- range $i, $f := .ParentFields}}{{if $i}}, {{end}}{{lower $f.Name}}{{end -}});{{end}}
        {{range .ExportFields}}
        this.{{.Name}} = {{lower .Name}};
        {{end}}
    }
}
