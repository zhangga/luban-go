package refs

type ITType interface {
	TryParseFrom(s string) bool
}

type TType struct {
	TypeName     string
	IsNullable   bool
	IsCollection bool
	IsBean       bool
	IsEnum       bool
	ElementType  ITType
	Tags         map[string]string
	Validators   []IDataValidator
}

func NewTType(isNullable bool, tags map[string]string) TType {
	return TType{
		IsNullable: isNullable,
		Tags:       tags,
	}
}

func (t *TType) HasTag(attrName string) bool {
	if len(t.Tags) == 0 {
		return false
	}
	_, ok := t.Tags[attrName]
	return ok
}

func (t *TType) GetTag(attrName string) (string, bool) {
	if len(t.Tags) == 0 {
		return "", false
	}
	v, ok := t.Tags[attrName]
	return v, ok
}

func (t *TType) GetTagOrDefault(attrName, defaultValue string) string {
	if len(t.Tags) == 0 {
		return defaultValue
	}
	v, ok := t.Tags[attrName]
	if !ok {
		return defaultValue
	}
	return v
}

type IDType interface {
}
