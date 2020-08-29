package godocgen

// PackageIndex contains all parsed types & functions information
type PackageIndex struct {
	Name string
	Dir  string
	Doc  string

	Consts   []Const
	Funcs    []Func
	Types    []Type
	Examples []Example
}

// Const is a single or a block of constants in a package
type Const struct {
	Doc     string
	Snippet string
}

// Func is a package level function
type Func struct {
	Name     string
	Doc      string
	Snippet  string
	Recv     string
	Examples []Example
}

// Type captures public types available at package level
type Type struct {
	Name     string
	Doc      string
	Snippet  string
	Consts   []Const
	Funcs    []Func
	Methods  []Func
	Examples []Example
}

// Example demonstrates use of functions & types in a package
type Example struct {
	Name    string
	Suffix  string
	Doc     string
	Snippet string
}
