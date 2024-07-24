package ast

type Node interface {
	Literal() string
	CSS(minifed bool) string
}

type Rule interface {
	Node
	parent() Rule
	rule()
}

type Declaration interface {
	Node
	declaration()
}
