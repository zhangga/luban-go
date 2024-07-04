package visitors

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/dataloader/excel"
)

var _ refs.ITypeVisitor1[*excel.ExcelStream, refs.DType] = (*ExcelStreamDataCreator)(nil)

type ExcelStreamDataCreator struct {
}

func (e *ExcelStreamDataCreator) Name() string {
	return "excel"
}

func (e *ExcelStreamDataCreator) AcceptBool(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptByte(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptShort(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptInt(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptLong(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptFloat(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptDouble(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptEnum(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptString(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptDateTime(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptBean(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptArray(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptList(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptSet(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (e *ExcelStreamDataCreator) AcceptMap(tType refs.TType, t *excel.ExcelStream) refs.DType {
	//TODO implement me
	panic("implement me")
}
