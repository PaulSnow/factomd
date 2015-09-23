// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/hex"
	"fmt"

	"github.com/FactomProject/factomd/btcd/wire"
	"github.com/FactomProject/factomd/common"
	fct "github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/database"
	"github.com/FactomProject/factomd/process"
)

var (
	db     database.Db
	inMsgQ chan wire.FtmInternalMsg
)

func ChainHead(chainid string) (*common.Hash, error) {
	h, err := atoh(chainid)
	if err != nil {
		return nil, err
	}
	c, err := db.FetchHeadMRByChainID(h)
	if err != nil {
		return nil, fmt.Errorf("Chain not found")
	}
	return c, nil
}

func CommitChain(c *common.CommitChain) error {
	m := wire.NewMsgCommitChain()
	m.CommitChain = c
	inMsgQ <- m
	return nil
}

func CommitEntry(c *common.CommitEntry) error {
	m := wire.NewMsgCommitEntry()
	m.CommitEntry = c
	inMsgQ <- m
	return nil
}

func FactoidTX(t fct.ITransaction) error {
	m := new(wire.MsgFactoidTX)
	m.SetTransaction(t)
	inMsgQ <- m
	return nil
}

func DBlockByKeyMR(keymr string) (*common.DirectoryBlock, error) {
	key, err := atoh(keymr)
	if err != nil {
		return nil, err
	}
	r, err := db.FetchDBlockByMR(key)
	if err != nil {
		return r, fmt.Errorf("DBlock not found")
	}
	return r, nil
}

func DBlockHead() (*common.DirectoryBlock, error) {
	_, height, err := db.FetchBlockHeightCache()
	if err != nil {
		return nil, err
	}
	block, err := db.FetchDBlockByHeight(uint32(height))
	if err != nil {
		return nil, err
	}
	block.BuildKeyMerkleRoot()
	return block, nil
}

func EBlockByKeyMR(keymr string) (*common.EBlock, error) {
	h, err := atoh(keymr)
	if err != nil {
		return nil, err
	}
	r, err := db.FetchEBlockByMR(h)
	if err != nil {
		return r, fmt.Errorf("EBlock not found")
	}
	return r, nil
}

func ECBalance(eckey string) (uint32, error) {
	key := new([32]byte)
	if p, err := hex.DecodeString(eckey); err != nil {
		return 0, err
	} else {
		copy(key[:], p)
	}
	val, _ := process.GetEntryCreditBalance(key)
	return uint32(val), nil
}

func EntryByHash(hash string) (*common.Entry, error) {
	h, err := atoh(hash)
	if err != nil {
		return nil, err
	}
	r, err := db.FetchEntryByHash(h)
	if err != nil {
		return r, err
	}
	if r == nil {
		return r, fmt.Errorf("Entry not found")
	}
	return r, nil
}

func RevealEntry(e *common.Entry) error {
	m := wire.NewMsgRevealEntry()
	m.Entry = e
	inMsgQ <- m
	return nil
}

func SetDB(d database.Db) {
	db = d
}

func SetInMsgQueue(q chan wire.FtmInternalMsg) {
	inMsgQ = q
}

func atoh(a string) (*common.Hash, error) {
	h := common.NewZeroHash()
	p, err := hex.DecodeString(a)
	if err != nil {
		return h, err
	}
	h.SetBytes(p)
	return h, nil
}