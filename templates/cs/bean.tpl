{{if .Namespace}}namespace {{.Namespace}}
{
{{end}}
    {{if .Comment}}/// <summary>
    /// {{.Comment}}
    /// </summary>{{end}}
    public {{if .IsDynamic}}abstract {{end}}class {{.Name}}{{if .Parent}} : {{.Parent}}{{end}}
    {
        {{range .ExportFields}}
        {{if .Comment}}/// <summary>
        /// {{.Comment}}
        /// </summary>{{end}}
        public readonly {{type .Type}} {{title .Name}};
        {{end}}
        
        public {{.Name}}(
            {{- range $i, $f := .HierarchyFields}}{{if $i}}, {{end}}{{type $f.Type}} {{lower $f.Name}}{{end -}}
        ){{if .Parent}} : base(
            {{- range $i, $f := .ParentFields}}{{if $i}}, {{end}}{{lower $f.Name}}{{end -}}
        ){{end}}
        {
            {{range .ExportFields}}
            this.{{title .Name}} = {{lower .Name}};
            {{end}}
        }
    }
{{if .Namespace}}
}
{{end}}