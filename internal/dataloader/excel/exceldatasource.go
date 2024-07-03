package excel

import (
	"fmt"
	"github.com/zhangga/luban/core/dataloader"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
)

var _ dataloader.IDataLoader = (*RowColumnDataSource)(nil)

type RowColumnDataSource struct {
	logger logger.Logger
	rawUrl string
	sheets []*RowColumnSheet
}

func NewDataSource(logger logger.Logger) dataloader.IDataLoader {
	return &RowColumnDataSource{
		logger: logger,
	}
}

func (ds *RowColumnDataSource) Names() []string {
	return []string{"xls", "xlsx", "xlsm", "xlm", "csv"}
}

func (ds *RowColumnDataSource) RawUrl() string {
	return ds.rawUrl
}

func (ds *RowColumnDataSource) Load(fileName, sheetName string) error {
	ds.logger.Debugf("load excel file: %s, sheet: %s", fileName, sheetName)
	ds.rawUrl = fileName

	for _, rawSheet := range loadRawSheets(fileName, sheetName) {
		sheet := NewRowColumnSheet(ds.rawUrl, sheetName, rawSheet.SheetName)
		sheet.Load(rawSheet)
		ds.sheets = append(ds.sheets, sheet)
	}

	if len(ds.sheets) == 0 {
		if len(sheetName) > 0 {
			return fmt.Errorf("excel: %s, sheet: %s, 不存在或者不是有效的单元簿(有效单元薄的A0单元格必须是##)", fileName, sheetName)
		} else {
			return fmt.Errorf("excel: %s, 不包含有效的单元薄(有效单元薄的A0单元格必须是##)", fileName)
		}
	}
	return nil
}

func (ds *RowColumnDataSource) ReadOne(ttype refs.TType) *refs.Record {
	panic("\"excel不支持单列读取模式")
}

func (ds *RowColumnDataSource) ReadMulti(ttype refs.TType) []*refs.Record {
	var datas []*refs.Record
	for _, sheet := range ds.sheets {
		for _, r := range sheet.Rows {
			row := r.Row
			tag := r.Tag
			if utils.IsIgnoreTag(tag) {
				continue
			}
			visitor := refs.MustGetTypeVisitor[refs.ITypeVisitor2[*RowColumnSheet, *TitleRow, refs.DType]]("sheet")
			data := refs.DispatchAccept2[*RowColumnSheet, *TitleRow, refs.DType](visitor, ttype, sheet, row)
			//data := ttype.ApplyVisitor2(visitor, sheet, row)
			datas = append(datas, &refs.Record{
				Data:   data,
				Source: sheet.UrlWithParams(),
				Tags:   utils.ParseTags(tag),
			})
		}
	}
	return datas
}
