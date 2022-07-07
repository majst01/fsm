//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM[string, string]
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Flows[string, string]{
			{Event: "open", Src: []string{"closed"}, Dst: "open"},
			{Event: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks[string, string]{
			"enter_state": func(e *fsm.CallbackReference[string, string]) { d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.CallbackReference[string, string]) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")

	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
