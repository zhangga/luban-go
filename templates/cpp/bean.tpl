#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>

{{if .Namespace}}
namespace {{.Namespace}} {
{{end}}

{{if .Comment}}
/**
 * {{.Comment}}
 */
{{end}}
class {{.Name}}{{if .Parent}} : public {{.Parent}}{{end}} {
public:
{{range .ExportFields}}
{{if .Comment}}    /**
     * {{.Comment}}
     */{{end}}
    {{type .Type}} {{.Name}};
{{end}}

    {{.Name}}({{- range $i, $f := .HierarchyFields}}{{if $i}}, {{end}}{{type $f.Type}} _{{.Name}}{{end -}})
    {{if .Parent}}: {{.Parent}}({{- range $i, $f := .ParentFields}}{{if $i}}, {{end}}_{{.Name}}{{end -}}){{end}} {
        {{range .ExportFields}}
        this->{{.Name}} = _{{.Name}};
        {{end}}
    }
    
    {{if .IsDynamic}}virtual ~{{.Name}}() = default;{{end}}
};

{{if .Namespace}}
}
{{end}}