package refs

import (
	"sync"
)

// TType types类型接口
type TType interface {
	TypeName() string // 类型名称
	IsBean() bool     // 是否bean类型
	IsNullable() bool
	HasTag(attrName string) bool
	GetTag(attrName string) (string, bool)
	GetTagOrDefault(attrName, defaultValue string) string

	TryParseFrom(s string) bool
}

// EmbedTType types类型的基类
type EmbedTType struct {
	IsNullable   bool              // 是否可空
	IsCollection bool              // 是否集合
	IsEnum       bool              // 是否枚举
	ElementType  TType             // 集合元素类型
	Tags         map[string]string // 标签
	Validators   []IDataValidator  // 验证器
}

func NewEmbedTType(isNullable bool, tags map[string]string) EmbedTType {
	return EmbedTType{
		IsNullable: isNullable,
		Tags:       tags,
	}
}

func (t *EmbedTType) HasTag(attrName string) bool {
	if len(t.Tags) == 0 {
		return false
	}
	_, ok := t.Tags[attrName]
	return ok
}

func (t *EmbedTType) GetTag(attrName string) (string, bool) {
	if len(t.Tags) == 0 {
		return "", false
	}
	v, ok := t.Tags[attrName]
	return v, ok
}

func (t *EmbedTType) GetTagOrDefault(attrName, defaultValue string) string {
	if len(t.Tags) == 0 {
		return defaultValue
	}
	v, ok := t.Tags[attrName]
	if !ok {
		return defaultValue
	}
	return v
}

type TypeCreator func(nullable bool, tags map[string]string, def IDefType, typ ...TType) TType

var (
	typeCreators = map[string]TypeCreator{}
	typeLocker   sync.RWMutex
)

func RegisterTypeCreator(creator TypeCreator) {
	// 创建实例
	instance := creator(false, nil, nil)
	typeLocker.Lock()
	defer typeLocker.Unlock()
	if _, ok := typeCreators[instance.TypeName()]; ok {
		panic("type creator: " + instance.TypeName() + " already registered")
	}
	typeCreators[instance.TypeName()] = creator
}

func GetTypeCreator(name string) TypeCreator {
	typeLocker.RLock()
	defer typeLocker.RUnlock()
	if creator, ok := typeCreators[name]; ok {
		return creator
	}
	return nil
}
