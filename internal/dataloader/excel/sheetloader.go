package excel

import (
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	metaSep                          = "#"
	maxEmptyRowCountOfInterruptParse = 60
)

var (
	knownSpecialTags = []string{"var", "+", "type", "desc", "comment", "column", "group"}
)

// LoadRawSheets 加载成RawSheet数据结构
func loadRawSheets(fileName, sheetName string) []*RawSheet {
	if ext := filepath.Ext(fileName); ext == ".csv" {
		return []*RawSheet{loadCsv(fileName)}
	} else {
		return loadExcelSheet(fileName, sheetName)
	}
}

// loadExcelSheet 加载excel表格
func loadExcelSheet(fileName, sheetName string) []*RawSheet {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		panic(fmt.Errorf("打开文件: %s, 错误: %s", fileName, err))
	}
	defer file.Close()

	var sheets []*RawSheet
	for _, name := range file.GetSheetList() {
		if len(sheetName) == 0 || name == sheetName {
			if rows, err := file.GetRows(name); err != nil {
				panic(fmt.Errorf("读取文件: %s, sheet: %s, 错误: %s", fileName, name, err))
			} else {
				sheet := parseRawSheet(newDataReader(rows, fileName, sheetName))
				sheets = append(sheets, sheet)
			}
		}
	}
	return sheets
}

// 加载csv文件
func loadCsv(fileName string) *RawSheet {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("打开文件: %s, 错误: %s", fileName, err))
	}
	defer file.Close()

	var rows [][]string
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("读取文件: %s, 错误: %s", fileName, err))
		} else {
			rows = append(rows, line)
		}
	}
	return parseRawSheet(newDataReader(rows, fileName, ""))
}

func parseRawSheet(reader *dataReader) *RawSheet {
	orientRow, tableName, ok := tryParseMeta(reader)
	if !ok {
		return nil
	}

	cells := parseRawSheetContent(reader, orientRow, false)
	validateTitles(cells, reader.fileName)
	title := parseTitle(reader, cells, orientRow)
	var dataCells [][]*Cell
	for _, row := range cells {
		if isNotDataRow(row) {
			continue
		}
		dataCells = append(dataCells, row)
	}
	return &RawSheet{
		Title:     title,
		TableName: tableName,
		SheetName: reader.sheetName,
		Cells:     dataCells,
	}
}

func parseRawSheetContent(reader *dataReader, orientRow bool, headerOnly bool) [][]*Cell {
	// TODO 优化性能
	// 几个思路
	// 1. 没有 title 的列不加载
	// 2. 空行优先跳过
	// 3. 跳过null或者empty的单元格
	var originRows [][]*Cell
	rowIndex, consecutiveEmptyRowCount := 0, 0
	for {
		var row []*Cell
		line := reader.ReadRow()
		if line == nil {
			break
		}
		for i, v := range line {
			row = append(row, NewCell(rowIndex, i, v, true))
		}
		originRows = append(originRows, row)
		if orientRow && headerOnly && !isHeaderRow(row) {
			break
		}
		rowIndex++
		if isEmptyRow(row) {
			consecutiveEmptyRowCount++
			if consecutiveEmptyRowCount > maxEmptyRowCountOfInterruptParse {
				logger.Errorf("excel: %s, sheet: %s 连续空行超过%d行，可能是数据错误，解析中断", reader.fileName, reader.sheetName, maxEmptyRowCountOfInterruptParse)
				break
			}
		} else {
			consecutiveEmptyRowCount = 0
		}
	}

	var finalRows [][]*Cell
	if orientRow {
		finalRows = originRows
	} else {
		// 转置这个行列
		maxColumn := 0
		for _, row := range originRows {
			if len(row) > maxColumn {
				maxColumn = len(row)
			}
		}
		for i := 0; i < maxColumn; i++ {
			var row []*Cell
			for j, originRow := range originRows {
				if i < len(originRow) {
					row = append(row, originRow[i])
				} else {
					row = append(row, NewCell(j, i, "", false))
				}
			}
			finalRows = append(finalRows, row)
		}
	}
	return finalRows
}

// parseTitle 解析标题
func parseTitle(reader *dataReader, cells [][]*Cell, orientRow bool) *Title {
	var maxColSize int
	for _, row := range cells {
		if len(row) > maxColSize {
			maxColSize = len(row)
		}
	}
	rootTitle := &Title{
		Root:      true,
		Name:      "__root__",
		Tags:      map[string]string{},
		FromIndex: 0,
		ToIndex:   maxColSize - 1,
	}
	topTitleRowIndex := tryFindTopTitle(cells)
	if topTitleRowIndex == -1 {
		panic(fmt.Errorf("文件: %s, sheet: %s, 没有定义任何有效 标题行", reader.fileName, reader.sheetName))
	}

	parseSubTitles(reader, rootTitle, cells, orientRow, topTitleRowIndex)
	rootTitle.Init()
	if !rootTitle.HasSubTitle() {
		panic(fmt.Errorf("文件: %s, sheet: %s, 没有定义任何有效列", reader.fileName, reader.sheetName))

	}
	return rootTitle
}

func parseSubTitles(reader *dataReader, rootTitle *Title, cells [][]*Cell, orientRow bool, topTitleRowIndex int) {
	titleRow := cells[topTitleRowIndex]
	for i := rootTitle.FromIndex; i <= rootTitle.ToIndex; i++ {
		nameAndAttrs := titleRow[i].ValueOrEmpty()
		if isIgnoreTitle(nameAndAttrs) {
			continue
		}

		titleName, tags := parseNameAndMetaAttrs(reader, nameAndAttrs)
		var subTitle *Title
		// [field,,,, field] 形成多列字段
		if strings.HasPrefix(titleName, "[") {
			startIndex := i
			titleName = titleName[1:]
			findEndPair := false
			for i++; i <= rootTitle.ToIndex; i++ {
				endNamePair := titleRow[i].ValueOrEmpty()
				if len(endNamePair) == 0 {
					continue
				}
				if !strings.HasSuffix(endNamePair, "]") || endNamePair[:len(endNamePair)-1] != titleName {
					panic(fmt.Errorf("excel: %s, sheet: %s, 列: '[%s' 后第一个有效列必须为匹配 '%s], 却发现为: %s' ", reader.fileName, reader.sheetName, titleName, titleName, endNamePair))
				}
				findEndPair = true
				break
			}
			if !findEndPair {
				panic(fmt.Errorf("excel: %s, sheet: %s, 列: '[%s' 后没有找到匹配的 '%s]' ", reader.fileName, reader.sheetName, titleName, titleName))
			}
			subTitle = &Title{
				Name:      titleName,
				Tags:      tags,
				FromIndex: startIndex,
				ToIndex:   i,
			}
		} else {
			if t, ok := rootTitle.SubTitles[titleName]; ok {
				if t.FromIndex != i {
					panic(fmt.Errorf("excel: %s, sheet: %s, 列: %s 重复", reader.fileName, reader.sheetName, titleName))
				} else {
					subTitle = t
					continue
				}
			}
			subTitle = &Title{
				Name:      titleName,
				Tags:      tags,
				FromIndex: i,
				ToIndex:   i,
			}
		}
		if topTitleRowIndex < len(cells) {
			nextRowIndex := tryFindNextSubFieldRowIndex(reader, cells, topTitleRowIndex)
			if nextRowIndex != -1 {
				parseSubTitles(reader, subTitle, cells, orientRow, nextRowIndex)
			}
		}
		rootTitle.AddSubTitle(subTitle)
	}
}

// tryFindTopTitle 尝试找到顶级标题，如果找不到则返回-1
func tryFindTopTitle(cells [][]*Cell) int {
	for i, row := range cells {
		if len(row) == 0 {
			break
		}
		rowTag := strings.ToLower(row[0].ValueOrEmpty())
		if !strings.HasPrefix(rowTag, "##") {
			break
		}
		if strings.Index(rowTag[2:], "&") != -1 {
			panic(fmt.Errorf("excel标题头不再使用'&'作为分割符，请改为: %s", metaSep))
		}

		var tags []string
		for _, s := range strings.Split(rowTag[2:], metaSep) {
			tag := strings.TrimSpace(s)
			if len(tag) == 0 {
				continue
			}
			tags = append(tags, tag)
		}
		if utils.Contain(tags, "field") || utils.Contain(tags, "var") || utils.Contain(tags, "+") {
			return i
		}

		// 出于历史兼容性，对第一行特殊处理，如果不包含任何tag或者只包含column，则也认为是标题行
		if i == 0 && (len(tags) == 0 || (len(tags) == 1 && tags[0] == "column")) {
			return i
		}
	}
	return -1
}

func tryFindNextSubFieldRowIndex(reader *dataReader, cells [][]*Cell, startRowIndex int) int {
	for i := startRowIndex + 1; i < len(cells); i++ {
		row := cells[i]
		if len(row) == 0 {
			break
		}
		rowTag := strings.ToLower(row[0].ValueOrEmpty())
		if rowTag == "##field" || rowTag == "##var" || rowTag == "##+" {
			return i
		} else if !strings.HasPrefix(rowTag, "##") {
			break
		}
	}
	return -1
}

func isHeaderRow(row []*Cell) bool {
	if len(row) == 0 {
		return false
	}
	value := row[0].ValueOrEmpty()
	if len(value) == 0 {
		return false
	}
	return strings.HasPrefix(value, "##")
}

func isEmptyRow(row []*Cell) bool {
	if len(row) == 0 {
		return true
	}
	for _, cell := range row {
		if len(cell.ValueOrEmpty()) > 0 {
			return false
		}
	}
	return true
}

func isNotDataRow(row []*Cell) bool {
	if len(row) == 0 {
		return true
	}
	s := row[0].ValueOrEmpty()
	return strings.HasPrefix(s, "##")
}

func tryParseMeta(reader *dataReader) (orientRow bool, tableName string, ok bool) {
	meta := strings.TrimSpace(reader.GetString(0, 0))
	// meta 行 必须以 ##为第一个单元格内容,紧接着 key:value 形式 表达meta属性
	if len(meta) == 0 || !strings.HasPrefix(meta, "##") {
		return true, "", false
	}
	if strings.Contains(meta, "&") {
		panic(fmt.Errorf("excel标题头不再使用'&'作为分割符，请改为: %s", metaSep))
	}

	orientRow = true
	for _, attr := range strings.Split(meta[2:], metaSep) {
		if len(attr) == 0 {
			continue
		}
		kvIndex := strings.Index(attr, "=")
		var key, value string
		if kvIndex == -1 {
			key = strings.TrimSpace(attr)
		} else {
			key = strings.TrimSpace(attr[:kvIndex])
			value = strings.TrimSpace(attr[kvIndex+1:])
		}
		switch key {
		case "field", "+", "var", "comment", "desc", "type":
		case "row":
			orientRow = true
		case "column":
			orientRow = false
		case "table":
			tableName = value
		default:
			panic(fmt.Errorf("非法单元薄 meta 属性定义: %s, 合法属性有: +,var,row,column,table=<tableName>", attr))
		}
	}
	return orientRow, tableName, true
}

func parseNameAndMetaAttrs(reader *dataReader, nameAndAttrs string) (titleName string, tags map[string]string) {
	if strings.Contains(nameAndAttrs, "&") {
		panic(fmt.Errorf("excel: %s, sheet: %s, 标题头不再使用'&'作为分割符，请改为: %s", reader.fileName, reader.sheetName, metaSep))
	}

	attrs := strings.Split(nameAndAttrs, metaSep)
	titleName = strings.TrimSpace(attrs[0])
	tags = make(map[string]string)
	// * 开头的表示是多行
	if strings.HasPrefix(titleName, "*") {
		titleName = titleName[1:]
		tags["multi_rows"] = "1"
	} else if strings.HasPrefix(titleName, "!") {
		titleName = titleName[1:]
		tags["non_empty"] = "1"
	}
	for i := 1; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], "=")
		if len(attrPair) != 2 {
			panic(fmt.Errorf("excel: %s, sheet: %s, 标题头属性定义错误: %s", reader.fileName, reader.sheetName, nameAndAttrs))
		}
		tags[strings.TrimSpace(attrPair[0])] = strings.TrimSpace(attrPair[1])
	}
	return titleName, tags
}

func validateTitles(rows [][]*Cell, fileName string) {
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		rowTag := row[0].ValueOrEmpty()
		if len(rowTag) == 0 {
			continue
		}
		if !strings.HasPrefix(rowTag, "##") {
			break
		}

		for _, s := range strings.Split(rowTag[2:], metaSep) {
			tag := strings.TrimSpace(s)
			if len(tag) == 0 {
				continue
			}
			if !utils.Contain(knownSpecialTags, tag) {
				logger.Fatalf("文件: %s, 行标签: %s, 包含未知tag: %s，是否有拼写错误?", fileName, rowTag, tag)
			}
		}
	}
}

func isIgnoreTitle(title string) bool {
	return len(title) == 0 || strings.HasPrefix(title, "#")
}
