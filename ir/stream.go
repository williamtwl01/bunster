package ir

import (
	"fmt"
)

const (
	FLAG_READ   = "STREAM_FLAG_READ"
	FLAG_WRITE  = "STREAM_FLAG_WRITE"
	FLAG_RW     = "STREAM_FLAG_RW"
	FLAG_APPEND = "STREAM_FLAG_APPEND"
)

type OpenStream struct {
	Name   string
	Target Instruction
	Mode   string
	Label  string
}

func (of OpenStream) togo() string {
	var ret = "return"
	if of.Label != "" {
		ret = "goto " + of.Label
	}

	return fmt.Sprintf(
		`%s, err := streamManager.OpenStream(%s, runtime.%s)
		if err != nil {
			shell.HandleError(err)
			%s
		}
		`, of.Name, of.Target.togo(), of.Mode, ret)
}

type NewBuffer struct {
	Value Instruction
}

func (b NewBuffer) togo() string {
	return fmt.Sprintf("runtime.NewBuffer(%s, false)", b.Value.togo())
}

type NewPipeBuffer struct {
	Value Instruction
	Name  string
	Label string
}

func (b NewPipeBuffer) togo() string {
	return fmt.Sprintf(
		`%s, err := runtime.NewBufferedStream(%s)
		if err != nil {
			shell.HandleError(err)
			return
		}
		`, b.Name, b.Value.togo())
}

type CloneStreamManager struct {
	DeferDestroy bool
}

func (c CloneStreamManager) togo() string {
	var deferDestroy string
	if c.DeferDestroy {
		deferDestroy = "defer streamManager.Destroy()\n"
	}
	return fmt.Sprintf("streamManager := streamManager.Clone() \n %s", deferDestroy)
}

type AddStream struct {
	Fd         string
	StreamName string
	Proxy      bool
}

func (as AddStream) togo() string {
	return fmt.Sprintf("streamManager.Add(`%s`, %s, %t)\n", as.Fd, as.StreamName, as.Proxy)
}

type SetStream struct {
	Name string
	Fd   Instruction
}

func (as SetStream) togo() string {
	return fmt.Sprintf(
		`if stream, err := streamManager.Get(%s); err != nil{
			shell.HandleError(err)
			return
		}else{
			%s = stream
		}
		`, as.Fd.togo(), as.Name)
}

type DuplicateStream struct {
	Old   string
	New   Instruction
	Label string
}

func (as DuplicateStream) togo() string {
	return fmt.Sprintf(
		`if err := streamManager.Duplicate("%s", %s); err != nil {
			shell.HandleError(err)
			return
		}
	`, as.Old, as.New.togo())
}

type CloseStream struct {
	Fd    Instruction
	Label string
}

func (c CloseStream) togo() string {
	return fmt.Sprintf(
		`if err := streamManager.Close(%s); err != nil {
			shell.HandleError(err)
			return
		}
	`, c.Fd.togo())
}
