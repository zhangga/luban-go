-- Lua 侧通常通过 require table 来使用，Bean 通常对应的是 table
-- {{.Name}} {{if .Comment}} {{.Comment}} {{end}}
local {{.Name}} = {}

{{if .IsDynamic}}
-- This is an abstract/dynamic class.
{{end}}
{{if .Parent}}
-- Inherits from {{.Parent}}
{{end}}

function {{.Name}}.New({{- range $i, $f := .HierarchyFields}}{{if $i}}, {{end}}{{lower $f.Name}}{{end -}})
    local obj = {}
    {{if .Parent}}
    -- inherit fields (assuming simple copy or metatable here, using simple assignment for generated struct)
    {{end}}
    {{range .HierarchyFields}}
    obj.{{.Name}} = {{lower .Name}} {{if .Comment}} -- {{.Comment}}{{end}}
    {{end}}
    return obj
end

return {{.Name}}