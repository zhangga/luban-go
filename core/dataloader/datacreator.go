package dataloader

import (
	"fmt"
	"strconv"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/types"
)

// DataCreator 负责将字符串或 DataStream 转换为对应 TType 的 DType 数据
type DataCreator struct{}

func NewDataCreator() *DataCreator {
	return &DataCreator{}
}

func (c *DataCreator) CreateData(t types.TType, stream *DataStream) (datas.DType, error) {
	str, ok := stream.ReadString()
	if !ok || str == "" {
		return nil, nil // 空值处理
	}

	switch ty := t.(type) {
	case *types.TBool:
		b, _ := strconv.ParseBool(str)
		return datas.NewDBool(b), nil
	case *types.TByte:
		n, _ := strconv.ParseUint(str, 10, 8)
		return datas.NewDByte(byte(n)), nil
	case *types.TShort:
		n, _ := strconv.ParseInt(str, 10, 16)
		return datas.NewDShort(int16(n)), nil
	case *types.TInt:
		n, _ := strconv.ParseInt(str, 10, 32)
		return datas.NewDInt(int32(n)), nil
	case *types.TLong:
		n, _ := strconv.ParseInt(str, 10, 64)
		return datas.NewDLong(int64(n)), nil
	case *types.TFloat:
		n, _ := strconv.ParseFloat(str, 32)
		return datas.NewDFloat(float32(n)), nil
	case *types.TDouble:
		n, _ := strconv.ParseFloat(str, 64)
		return datas.NewDDouble(float64(n)), nil
	case *types.TString:
		return datas.NewDString(str), nil
	case *types.TEnum:
		// 枚举暂以 Int 形式兜底，也可以是 String
		if n, err := strconv.ParseInt(str, 10, 32); err == nil {
			return datas.NewDEnum(int32(n), str, ty), nil
		}
		return datas.NewDEnum(0, str, ty), nil
	case *types.TArray:
		// 解析嵌套结构，对于容器，如果只有一格文本，再切分一次，或是直接从后续 token 获取
		// 这里简化：如果传来的是单个包含多个元素的串，我们在外部需要用新的 stream
		return nil, fmt.Errorf("array creation should be handled in list logic")
	default:
		return datas.NewDString(str), nil
	}
}

// CreateList 解析 List 类型的单元格，默认按照逗号或者分号分割元素
func (c *DataCreator) CreateList(t *types.TList, cellStr string) (datas.DType, error) {
	if cellStr == "" {
		return datas.NewDList(t, []datas.DType{}), nil
	}
	sep := t.GetTagOrDefault("sep", ",")
	stream := NewDataStream(cellStr, sep)

	elements := make([]datas.DType, 0)
	elemType := t.ElementType()

	for stream.HasNext() {
		elem, err := c.CreateData(elemType, stream)
		if err != nil {
			return nil, err
		}
		if elem != nil {
			elements = append(elements, elem)
		}
	}
	return datas.NewDList(t, elements), nil
}

// CreateMap 解析 Map 类型的单元格，默认 k:v,k:v 或 k,v,k,v
func (c *DataCreator) CreateMap(t *types.TMap, cellStr string) (datas.DType, error) {
	if cellStr == "" {
		return datas.NewDMap(t, map[datas.DType]datas.DType{}), nil
	}
	sep := t.GetTagOrDefault("sep", ",")
	stream := NewDataStream(cellStr, sep)

	dataMap := make(map[datas.DType]datas.DType)
	kType := t.KeyType()
	vType := t.ValueType()

	for stream.HasNext() {
		// Key
		k, err := c.CreateData(kType, stream)
		if err != nil {
			return nil, err
		}

		// 有时候 k:v 是用冒号连在一起的，这里我们假设 sep 分割了平铺的 k,v,k,v
		// 真实的 Luban 中通常可以通过 map 的 sep 定义来决定。
		if !stream.HasNext() {
			break
		}

		// Value
		v, err := c.CreateData(vType, stream)
		if err != nil {
			return nil, err
		}

		if k != nil {
			dataMap[k] = v
		}
	}

	return datas.NewDMap(t, dataMap), nil
}

// CreateField 解析并创建任意字段
func (c *DataCreator) CreateField(t types.TType, cellStr string) (datas.DType, error) {
	if cellStr == "" {
		return nil, nil
	}

	switch ty := t.(type) {
	case *types.TList:
		return c.CreateList(ty, cellStr)
	case *types.TMap:
		return c.CreateMap(ty, cellStr)
	case *types.TArray:
		// Array 的解析复用 List 的逻辑
		listType := types.NewTList(ty.IsNullable(), ty.ElementType(), ty.Tags())
		l, err := c.CreateList(listType, cellStr)
		if err != nil {
			return nil, err
		}
		if dList, ok := l.(*datas.DList); ok {
			return datas.NewDArray(ty, dList.Elements), nil
		}
		return nil, fmt.Errorf("array creation error")
	default:
		// 基础类型直接走 stream
		stream := NewDataStream(cellStr, ",")
		return c.CreateData(t, stream)
	}
}
