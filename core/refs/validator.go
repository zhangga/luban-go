package refs

type IDataValidator interface {
	Args() string
	Compile(owner IDefField, ttype ITType)
	Validate(ttype ITType, dtype IDType)
}
