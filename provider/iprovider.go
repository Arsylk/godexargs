package provider

type ClassDefId = uint32

type ClassDef struct {
	Id    ClassDefId
	Name  string
	Flags uint32
}

type MethodDefId = uint32

type MethodParamDef struct {
	Name string
	Type uint8
}

type MethodDef struct {
	Id     MethodDefId
	Name   string
	Flags  uint32
	Params []MethodParamDef
	Ret    string
}

type IProvider interface {
	ListClasses() *[]ClassDef
	ListMethods(id ClassDefId) *[]MethodDef
}

type ProviderConstructor[T IProvider] func(...any) *T
