package visitors

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/dataloader/excel"
	"github.com/zhangga/luban/internal/types"
)

var _ refs.ITypeVisitor2[*excel.RowColumnSheet, *excel.TitleRow, refs.DType] = (*SheetDataCreator)(nil)

type SheetDataCreator struct {
}

func (d *SheetDataCreator) Name() string {
	return "sheet"
}

func (d *SheetDataCreator) AcceptBool(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptByte(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptShort(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptInt(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptLong(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptFloat(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptDouble(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptEnum(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptString(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptDateTime(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptBean(ttype refs.TType, sheet *excel.RowColumnSheet, row *excel.TitleRow) refs.DType {
	sep := row.SelfTitle.Sep
	if row.Row != nil {
		stream := row.AsStream(sep)
		if ttype.IsNullable() && stream.TryReadEOF() {
			return nil
		}
		visitor := refs.MustGetTypeVisitor[refs.ITypeVisitor1[*excel.ExcelStream, refs.DType]]("excel")
		return refs.DispatchAccept1[*excel.ExcelStream, refs.DType](visitor, ttype, stream)
	}

	if row.Rows != nil {
		stream := row.AsMultiRowConcatStream(sep)
		if ttype.IsNullable() && stream.TryReadEOF() {
			return nil
		}
		visitor := refs.MustGetTypeVisitor[refs.ITypeVisitor1[*excel.ExcelStream, refs.DType]]("excel")
		return refs.DispatchAccept1[*excel.ExcelStream, refs.DType](visitor, ttype, stream)
	}

	if row.Fields != nil {
		tbean := ttype.(*types.TBean)
		originBean := tbean.DefBean
		sep += originBean.Sep
		if originBean.IsAbstractType() {
			panic("implement me")
			//typeTitle := row.get
		}
	}

	if row.Elements != nil {
		stream := row.AsMultiRowConcatElements(sep)
		visitor := refs.MustGetTypeVisitor[refs.ITypeVisitor1[*excel.ExcelStream, refs.DType]]("excel")
		return refs.DispatchAccept1[*excel.ExcelStream, refs.DType](visitor, ttype, stream)
	}

	panic("unsupported bean type")
}

func (d *SheetDataCreator) AcceptArray(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptList(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptSet(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}

func (d *SheetDataCreator) AcceptMap(ttype refs.TType, t *excel.RowColumnSheet, s *excel.TitleRow) refs.DType {
	//TODO implement me
	panic("implement me")
}
