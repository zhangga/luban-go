package datatarget

import (
	"encoding/binary"
	"math"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
)

// ByteBuf 用于存储二进制数据
type ByteBuf struct {
	data []byte
}

func NewByteBuf() *ByteBuf {
	return &ByteBuf{data: make([]byte, 0)}
}

func (b *ByteBuf) WriteBool(v bool) {
	if v {
		b.data = append(b.data, 1)
	} else {
		b.data = append(b.data, 0)
	}
}

func (b *ByteBuf) WriteByte(v byte) {
	b.data = append(b.data, v)
}

func (b *ByteBuf) WriteShort(v int16) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(v))
	b.data = append(b.data, buf...)
}

func (b *ByteBuf) WriteInt(v int32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(v))
	b.data = append(b.data, buf...)
}

func (b *ByteBuf) WriteLong(v int64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(v))
	b.data = append(b.data, buf...)
}

func (b *ByteBuf) WriteFloat(v float32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, math.Float32bits(v))
	b.data = append(b.data, buf...)
}

func (b *ByteBuf) WriteDouble(v float64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(v))
	b.data = append(b.data, buf...)
}

func (b *ByteBuf) WriteString(v string) {
	bytes := []byte(v)
	// 原版 luban 中通常有 Size 的写入，这里简单使用 int32 作为长度
	b.WriteInt(int32(len(bytes)))
	b.data = append(b.data, bytes...)
}

func (b *ByteBuf) Bytes() []byte {
	return b.data
}

// BinDataVisitor 是一个实现 IDataVisitor 接口的访问者，将 DType 转换为二进制
type BinDataVisitor struct {
	TargetGroups []string
	buf          *ByteBuf
}

func NewBinDataVisitor(targetGroups []string) *BinDataVisitor {
	return &BinDataVisitor{
		TargetGroups: targetGroups,
		buf:          NewByteBuf(),
	}
}

func (v *BinDataVisitor) Buf() *ByteBuf {
	return v.buf
}

func (v *BinDataVisitor) VisitDBool(d *datas.DBool) interface{} {
	v.buf.WriteBool(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDByte(d *datas.DByte) interface{} {
	v.buf.WriteByte(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDShort(d *datas.DShort) interface{} {
	v.buf.WriteShort(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDInt(d *datas.DInt) interface{} {
	v.buf.WriteInt(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDLong(d *datas.DLong) interface{} {
	v.buf.WriteLong(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDFloat(d *datas.DFloat) interface{} {
	v.buf.WriteFloat(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDDouble(d *datas.DDouble) interface{} {
	v.buf.WriteDouble(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDString(d *datas.DString) interface{} {
	v.buf.WriteString(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDEnum(d *datas.DEnum) interface{} {
	v.buf.WriteInt(d.Value)
	return nil
}

func (v *BinDataVisitor) VisitDBean(d *datas.DBean) interface{} {
	// 如果是多态，需要写入类型 id (简单起见，可以写类型名字的字符串，或者对应的 typeid。由于目前没有 typeid 的哈希计算，这里写入类型名字)
	if d.Type.IsDynamic() && d.ImplType != nil {
		v.buf.WriteString(d.ImplType.DefBean.Name())
	}

	implBean := d.Type
	if d.ImplType != nil {
		implBean = d.ImplType
	}

	if implBean != nil && implBean.DefBean != nil {
		if beanImpl, ok := implBean.DefBean.(*defs.DefBeanImpl); ok {
			for i, field := range beanImpl.HierarchyFields {
				// 添加 group 过滤
				if !field.NeedExport(v.TargetGroups) {
					continue
				}

				if i < len(d.Fields) {
					if d.Fields[i] != nil {
						d.Fields[i].Accept(v)
					}
				}
			}
		}
	}
	return nil
}

func (v *BinDataVisitor) VisitDArray(d *datas.DArray) interface{} {
	v.buf.WriteInt(int32(len(d.Elements)))
	for _, e := range d.Elements {
		if e != nil {
			e.Accept(v)
		}
	}
	return nil
}

func (v *BinDataVisitor) VisitDList(d *datas.DList) interface{} {
	v.buf.WriteInt(int32(len(d.Elements)))
	for _, e := range d.Elements {
		if e != nil {
			e.Accept(v)
		}
	}
	return nil
}

func (v *BinDataVisitor) VisitDSet(d *datas.DSet) interface{} {
	v.buf.WriteInt(int32(len(d.Elements)))
	for _, e := range d.Elements {
		if e != nil {
			e.Accept(v)
		}
	}
	return nil
}

func (v *BinDataVisitor) VisitDMap(d *datas.DMap) interface{} {
	v.buf.WriteInt(int32(len(d.Datas)))
	for k, val := range d.Datas {
		k.Accept(v)
		if val != nil {
			val.Accept(v)
		}
	}
	return nil
}

func (v *BinDataVisitor) VisitDDateTime(d *datas.DDateTime) interface{} {
	// 将 DateTime 作为 string 写入
	v.buf.WriteString(d.String())
	return nil
}