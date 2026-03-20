syntax = "proto3";
{{if .Namespace}}
package {{.Namespace}};
{{end}}

{{if .Comment}}// {{.Comment}}{{end}}
message {{.Name}} {
    {{range $i, $f := .ExportFields}}
    {{if $f.Comment}}// {{$f.Comment}}{{end}}
    {{type $f.Type}} {{lower $f.Name}} = {{add $i 1}};
    {{end}}
}
