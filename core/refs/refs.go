package refs

// NoUnkeyedLiterals can be embedded in a struct to prevent unkeyed literals.
type NoUnkeyedLiterals struct{}

type UnimplementedTable interface {
	MustEmbedUnimplementedTable()
}

type UnimplementedBean interface {
	MustEmbedUnimplementedBean()
}

type UnimplementedEnum interface {
	MustEmbedUnimplementedEnum()
}

type UnimplementedRefGroup interface {
	MustEmbedUnimplementedRefGroup()
}
