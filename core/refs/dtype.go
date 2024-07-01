package refs

type DType interface {
	CompareTo(other DType) int
	String() string
	TypeName() string
}

type EmbedDType[T any] struct {
	Value T
}
