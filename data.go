package ast_base

var Result = &metadata{} // Result is the ast parse result

type metadata struct { // Apis Project global API cache
	Apis map[string][]*MethodInfo
}

// MethodInfo Api method info
type MethodInfo struct {
	Method  string // API method。such as: POST、GET、DELETE、PUT、OPTIONS、PATCH、HEAD
	APIPath string // API path
	Name    string // Func name
}
