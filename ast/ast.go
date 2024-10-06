package ast

type Node interface {
	node() // Just to distinguish it.
}

type Script struct {
	Statements []Node
}

type LogicalCommand struct {
	Left     Node
	Operator string // || or &&
	Right    Node
}

func (LogicalCommand) node() {}

type BackgroundCommand struct {
	Node
}

type PipelineCommand struct {
	Stderr  bool
	Command Command
}

type Pipeline []PipelineCommand

func (Pipeline) node() {}

type Word string

func (Word) node() {}

type Redirection struct {
	Src    Node
	Method string
	Dst    Node
}

func (Redirection) node() {}

type FileDescriptor string

func (FileDescriptor) node() {}

type Command struct {
	Name         Node
	Args         []Node
	Redirections []Redirection
}

func (Command) node() {}

type SimpleExpansion string

func (SimpleExpansion) node() {}

type Concatination struct {
	Nodes []Node
}

func (Concatination) node() {}
