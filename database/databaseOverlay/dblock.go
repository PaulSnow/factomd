// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package databaseOverlay

import (
	"github.com/FactomProject/factomd/common/interfaces"
)

// ProcessDBlockBatche inserts the DBlock and update all it's dbentries in DB
func (db *Overlay) ProcessDBlockBatch(dblock interfaces.DatabaseBatchable) error {
	return db.ProcessBlockBatch([]byte{byte(TBL_DB)}, []byte{byte(TBL_DB_NUM)}, []byte{byte(TBL_DB_MR)}, dblock)
}

// FetchHeightRange looks up a range of blocks by the start and ending
// heights.  Fetch is inclusive of the start height and exclusive of the
// ending height. To fetch all hashes from the start height until no
// more are present, use the special id `AllShas'.
func (db *Overlay) FetchHeightRange(startHeight, endHeight int64) ([]interfaces.IHash, error) {
	return db.FetchBlockIndexesInHeightRange([]byte{byte(TBL_DB_NUM)}, startHeight, endHeight)
}

// FetchBlockHeightBySha returns the block height for the given hash.  This is
// part of the database.Db interface implementation.
func (db *Overlay) FetchBlockHeightByKeyMR(sha interfaces.IHash, dst interfaces.DatabaseBatchable) (int64, error) {
	dblk, err := db.FetchDBlockByKeyMR(sha, dst)
	if err != nil {
		return -1, err
	}

	var height int64 = -1
	if dblk != nil {
		height = int64(dblk.GetDatabaseHeight())
	}

	return height, nil
}

// FetchDBlock gets an entry by hash from the database.
func (db *Overlay) FetchDBlockByKeyMR(keyMR interfaces.IHash, dst interfaces.DatabaseBatchable) (interfaces.DatabaseBatchable, error) {
	return db.FetchBlock([]byte{byte(TBL_DB)}, keyMR, dst)
}

// FetchDBlockByHeight gets an directory block by height from the database.
func (db *Overlay) FetchDBlockByHeight(dBlockHeight uint32, dst interfaces.DatabaseBatchable) (interfaces.DatabaseBatchable, error) {
	return db.FetchBlockByHeight([]byte{byte(TBL_DB_NUM)}, []byte{byte(TBL_DB_MR)}, dBlockHeight, dst)
}

// FetchDBHashByHeight gets a dBlockHash from the database.
func (db *Overlay) FetchDBHashByHeight(dBlockHeight uint32) (interfaces.IHash, error) {
	return db.FetchBlockIndexByHeight([]byte{byte(TBL_DB_NUM)}, dBlockHeight)
}

// FetchDBHashByMR gets a DBHash by MR from the database.
func (db *Overlay) FetchDBHashByMR(dBMR interfaces.IHash) (interfaces.IHash, error) {
	return db.FetchPrimaryIndexBySecondaryIndex([]byte{byte(TBL_DB_MR)}, dBMR)
}

// FetchDBlockByMR gets a directory block by merkle root from the database.
func (db *Overlay) FetchDBlockByHash(dBMR interfaces.IHash, dst interfaces.DatabaseBatchable) (interfaces.DatabaseBatchable, error) {
	return db.FetchBlockBySecondaryIndex([]byte{byte(TBL_DB_MR)}, []byte{byte(TBL_DB)}, dBMR, dst)
}

// FetchAllDBlocks gets all of the fbInfo
func (db *Overlay) FetchAllDBlocks(sample interfaces.BinaryMarshallableAndCopyable) ([]interfaces.BinaryMarshallableAndCopyable, error) {
	return db.FetchAllBlocksFromBucket([]byte{byte(TBL_DB)}, sample)
}