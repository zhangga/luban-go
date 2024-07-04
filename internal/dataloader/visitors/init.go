package visitors

import "github.com/zhangga/luban/core/refs"

// init 注册visitor
func init() {
	refs.RegisterTypeVisitor[*SheetDataCreator]()
	refs.RegisterTypeVisitor[*ExcelStreamDataCreator]()
}
