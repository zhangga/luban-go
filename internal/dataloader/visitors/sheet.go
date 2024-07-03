package visitors

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/dataloader/excel"
)

var _ refs.ITypeVisitor2[*excel.RowColumnSheet, *excel.TitleRow, refs.DType] = (*SheetDataCreator)(nil)

type SheetDataCreator struct {
}

func (d *SheetDataCreator) Name() string {
	return "sheet"
}

func (d *SheetDataCreator) AcceptBool(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptByte(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptShort(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptInt(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptLong(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptFloat(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptDouble(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptEnum(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptString(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptDateTime(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptBean(sheet *excel.RowColumnSheet, row *excel.TitleRow) refs.DType {
	sep := row.SelfTitle.Sep
	if row.Row != nil {

	}
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptArray(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptList(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptSet(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptMap(t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}
