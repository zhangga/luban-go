package refs

type IDataValidator interface {
	Args() string
	Compile(owner IDefField, ttype TType)
	Validate(ttype TType, dtype DType)
}
