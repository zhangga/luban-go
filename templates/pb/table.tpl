syntax = "proto3";
{{if .Namespace}}
package {{.Namespace}};
{{end}}

// import "{{.ValueType}}.proto"; // 真实生成中可能需要处理 import 依赖

message {{.Name}} {
    repeated {{.ValueType}} data_list = 1;
}
