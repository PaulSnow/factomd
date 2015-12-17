// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	. "github.com/FactomProject/factomd/common/entryBlock"
	. "github.com/FactomProject/factomd/common/interfaces"
	"io"
)

// MsgRevealChain implements the Message interface and represents a factom
// Reveal-Chain message.  It is used by client to reveal the chain.
type MsgRevealChain struct {
	FirstEntry IEBEntry
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgRevealChain) BtcEncode(w io.Writer, pver uint32) error {

	//FirstEntry
	bytes, err := msg.FirstEntry.MarshalBinary()
	if err != nil {
		return err
	}

	err = writeVarBytes(w, pver, bytes)
	if err != nil {
		return err
	}

	return nil
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgRevealChain) BtcDecode(r io.Reader, pver uint32) error {
	//FirstEntry
	bytes, err := readVarBytes(r, pver, MaxAppMsgPayload, CmdRevealChain)
	if err != nil {
		return err
	}

	msg.FirstEntry = new(Entry)
	err = msg.FirstEntry.UnmarshalBinary(bytes)
	if err != nil {
		return err
	}

	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgRevealChain) Command() string {
	return CmdRevealChain
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgRevealChain) MaxPayloadLength(pver uint32) uint32 {
	return MaxAppMsgPayload
}

// NewMsgInv returns a new bitcoin inv message that conforms to the Message
// interface.  See MsgInv for details.
func NewMsgRevealChain() *MsgRevealChain {
	return &MsgRevealChain{}
}

// Create a sha hash from the message binary (output of BtcEncode)
func (msg *MsgRevealChain) Sha() (IHash, error) {

	buf := bytes.NewBuffer(nil)
	msg.BtcEncode(buf, ProtocolVersion)
	var sha IHash
	_ = sha.SetBytes(Sha256(buf.Bytes()))

	return sha, nil
}