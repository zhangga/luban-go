package datas

// DType 定义了 Luban 运行时数据的基本接口
type DType interface {
	TypeName() string
	// Accept 访问者入口
	Accept(visitor IDataVisitor) interface{}
	String() string
}
