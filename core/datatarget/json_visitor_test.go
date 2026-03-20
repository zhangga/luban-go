package datatarget

import (
	"encoding/json"
	"testing"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestJsonVisitor_GroupFilter(t *testing.T) {
	// 创建带 Group 的 Bean 定义
	rawBean := &rawdefs.RawBean{
		Name: "TestBean",
		Fields: []*rawdefs.RawField{
			{Name: "CommonField", Type: "int"},
			{Name: "ClientField", Type: "string", Groups: []string{"client"}},
			{Name: "ServerField", Type: "string", Groups: []string{"server"}},
		},
	}
	
	defBean := defs.NewDefBeanImpl(rawBean)
	for _, f := range rawBean.Fields {
		defBean.Fields = append(defBean.Fields, defs.NewDefField(f))
		defBean.HierarchyFields = append(defBean.HierarchyFields, defs.NewDefField(f))
	}

	tBean := types.NewTBean(false, defBean, nil)

	// 创建数据
	dBean := datas.NewDBean(tBean, tBean, []datas.DType{
		datas.NewDInt(1),
		datas.NewDString("for client"),
		datas.NewDString("for server"),
	})

	// 测试 client group
	visitorClient := NewToJsonVisitor([]string{"client"})
	resClient := dBean.Accept(visitorClient)
	bytesClient, _ := json.Marshal(resClient)
	strClient := string(bytesClient)

	if !stringsContains(strClient, "CommonField") || !stringsContains(strClient, "ClientField") {
		t.Errorf("Client group should contain CommonField and ClientField, got %s", strClient)
	}
	if stringsContains(strClient, "ServerField") {
		t.Errorf("Client group should NOT contain ServerField, got %s", strClient)
	}

	// 测试 server group
	visitorServer := NewToJsonVisitor([]string{"server"})
	resServer := dBean.Accept(visitorServer)
	bytesServer, _ := json.Marshal(resServer)
	strServer := string(bytesServer)

	if !stringsContains(strServer, "CommonField") || !stringsContains(strServer, "ServerField") {
		t.Errorf("Server group should contain CommonField and ServerField, got %s", strServer)
	}
	if stringsContains(strServer, "ClientField") {
		t.Errorf("Server group should NOT contain ClientField, got %s", strServer)
	}
}

func stringsContains(s, substr string) bool {
	return len(s) >= len(substr) && func() bool {
		for i := 0; i <= len(s)-len(substr); i++ {
			if s[i:i+len(substr)] == substr {
				return true
			}
		}
		return false
	}()
}
