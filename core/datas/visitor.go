package datas

type IDataVisitor interface {
	VisitDBool(d *DBool) interface{}
	VisitDByte(d *DByte) interface{}
	VisitDShort(d *DShort) interface{}
	VisitDInt(d *DInt) interface{}
	VisitDLong(d *DLong) interface{}
	VisitDFloat(d *DFloat) interface{}
	VisitDDouble(d *DDouble) interface{}
	VisitDString(d *DString) interface{}
	VisitDText(d *DText) interface{}
	VisitDEnum(d *DEnum) interface{}
	VisitDBean(d *DBean) interface{}
	VisitDArray(d *DArray) interface{}
	VisitDList(d *DList) interface{}
	VisitDSet(d *DSet) interface{}
	VisitDMap(d *DMap) interface{}
	VisitDDateTime(d *DDateTime) interface{}
}
