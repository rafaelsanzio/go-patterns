package interfaceprototype

type INode interface {
	Print(string)
	Clone() INode
}
