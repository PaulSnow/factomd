package factom2

import (
	"fmt"
	"github.com/PaulSnow/factom2d/common/interfaces"
)

type Instance struct {
	Identity int // Stop-gap Identity index

	InMsg         chan interfaces.IMsg // Network messages come in through this channel
	NetOutMsg     chan interfaces.IMsg // Broadcast and p2p messages go out to the network from this channel
	NetOutInvalid chan interfaces.IMsg // Invalid messages are reported to the network over this channel
	APIQueue      chan interfaces.IMsg // API submits transactions over this channel
}

// Stub for node status
func (ins *Instance) Status() string {
	str := fmt.Sprintf(" F2Node%2d DBHeight %8d", ins.Identity, 0)
	return str
}

func (ins *Instance) Run() {
	for {
		msg := <-ins.InMsg
		fmt.Println(msg.String())
	}
}
