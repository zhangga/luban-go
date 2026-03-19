package types

// ITypeVisitor 定义了对 TType 进行访问的访问者接口
// 为了简化，在 Go 中我们使用一个接受具体类型的方法列表
type ITypeVisitor interface {
	VisitTBool(t *TBool) interface{}
	VisitTByte(t *TByte) interface{}
	VisitTShort(t *TShort) interface{}
	VisitTInt(t *TInt) interface{}
	VisitTLong(t *TLong) interface{}
	VisitTFloat(t *TFloat) interface{}
	VisitTDouble(t *TDouble) interface{}
	VisitTString(t *TString) interface{}
	VisitTEnum(t *TEnum) interface{}
	VisitTBean(t *TBean) interface{}
	VisitTArray(t *TArray) interface{}
	VisitTList(t *TList) interface{}
	VisitTSet(t *TSet) interface{}
	VisitTMap(t *TMap) interface{}
	VisitTDateTime(t *TDateTime) interface{}
}
