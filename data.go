package ast_base

type metadata struct { // Apis Project global API cache
	Apis map[string][]*MethodInfo

	// Annotations Project global API annotations cache
	Annotations map[string]map[string]string
}

// MethodInfo Api method info
type MethodInfo struct {
	Method  string // API method。such as: POST、GET、DELETE、PUT、OPTIONS、PATCH、HEAD
	APIPath string // API path
	Name    string // Func name
}

var Metadata = &metadata{}
