{{if .Namespace}}namespace {{.Namespace}}
{
{{end}}
    {{if .Comment}}/// <summary>
    /// {{.Comment}}
    /// </summary>{{end}}
    public enum {{.Name}}
    {
        {{range .Items}}
        {{if .Alias}}/// <summary>
        /// {{.Alias}}
        /// </summary>{{end}}
        {{.Name}}{{if .Value}} = {{.Value}}{{end}},
        {{end}}
    }
{{if .Namespace}}
}
{{end}}