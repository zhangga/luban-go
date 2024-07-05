package visitors

import (
	"fmt"
	"github.com/zhangga/luban/core/dataloader"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/dataloader/excel"
	"github.com/zhangga/luban/internal/datas"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/types"
	"strings"
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
			var typeTitle *excel.TitleRow
			if typeTitle = row.GetSubTitleNamedRow(dataloader.ExcelTypeNameKey); typeTitle == nil {
				typeTitle = row.GetSubTitleNamedRow(dataloader.FallbackTypeNameKey)
			}
			if typeTitle == nil {
				panic(fmt.Errorf("type: %s, 是多态类型,需要定义`%s`列来指定具体子类型", originBean.FullName(), dataloader.ExcelTypeNameKey))
			}
			valueTitle := row.GetSubTitleNamedRow(dataloader.ExcelValueNameKey)
			sep += tbean.GetTagOrDefault("sep", "")
			subType := strings.TrimSpace(typeTitle.Current())
			if len(subType) == 0 || subType == dataloader.BeanNullType {
				if !ttype.IsNullable() {
					panic(fmt.Errorf("type: %s, 不是可空类型, `%s`不能为空", originBean.FullName(), tbean.DefBean.FullName()))
				}
				return nil
			}

			implType := defs.GetImplTypeByNameOrAlias(originBean, subType)
			if valueTitle == nil {
				return datas.NewDBean(tbean, implType, d.CreateBeanFields(implType, sheet, row))
			}

			sep += valueTitle.SelfTitle.Sep
			if valueTitle.Row != nil {
				stream := valueTitle.AsStream(sep)
				if ttype.IsNullable() && stream.TryReadEOF() {
					return nil
				}
				return datas.NewDBean(tbean, implType, d.CreateBeanFieldsByStream(implType, stream))
			}
			if valueTitle.Rows != nil {
				stream := valueTitle.AsMultiRowConcatStream(sep)
				if ttype.IsNullable() && stream.TryReadEOF() {
					return nil
				}
				return datas.NewDBean(tbean, implType, d.CreateBeanFieldsByStream(implType, stream))
			}
			panic("unsupported abstract bean type")
		}

		if tbean.IsNullable() {
			var typeTitle *excel.TitleRow
			if typeTitle = row.GetSubTitleNamedRow(dataloader.ExcelTypeNameKey); typeTitle == nil {
				typeTitle = row.GetSubTitleNamedRow(dataloader.FallbackTypeNameKey)
			}
			if typeTitle == nil {
				panic(fmt.Errorf("type: %s, 是可空类型,需要定义`%s`列来指明是否可空", originBean.FullName(), dataloader.ExcelTypeNameKey))
			}
			subType := strings.TrimSpace(typeTitle.Current())
			if len(subType) == 0 || subType == dataloader.BeanNullType {
				return nil
			}
			if subType != dataloader.BeanNotNullType && subType != originBean.Name {
				panic(fmt.Errorf("type: %s, 可空标识: %s, 不合法（只能为`%s`或者`%s`或者`%s`)", originBean.FullName(), subType, dataloader.BeanNullType, dataloader.BeanNotNullType, originBean.Name))
			}
		}

		return datas.NewDBean(tbean, originBean, d.CreateBeanFields(originBean, sheet, row))
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

func (d *SheetDataCreator) CreateBeanFields(defBean *defs.DefBean, sheet *excel.RowColumnSheet, row *excel.TitleRow) []refs.DType {
	var fields []refs.DType
	for _, f := range defBean.HierarchyFields {
		fname := f.Name()
		fieldRow := row.GetSubTitleNamedRow(fname)
		if fieldRow == nil {
			panic(fmt.Errorf("bean: %s, 缺失列: %s, 请检查是否写错或者遗漏", defBean.FullName(), fname))
		}
		field := refs.DispatchAccept2[*excel.RowColumnSheet, *excel.TitleRow, refs.DType](d, f.CType, sheet, fieldRow)
		fields = append(fields, field)
	}
	return fields
}

func (d *SheetDataCreator) CreateBeanFieldsByStream(defBean *defs.DefBean, stream *excel.ExcelStream) []refs.DType {
	var fields []refs.DType
	for _, f := range defBean.HierarchyFields {
		visitor := refs.MustGetTypeVisitor[refs.ITypeVisitor1[*excel.ExcelStream, refs.DType]]("excel")
		field := refs.DispatchAccept1[*excel.ExcelStream, refs.DType](visitor, f.CType, stream)
		fields = append(fields, field)
	}
	return fields
}
