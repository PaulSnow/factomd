// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package database

import (
	. "github.com/FactomProject/factomd/common/constants"
	. "github.com/FactomProject/factomd/common/interfaces"
)

type FDatabase struct {
	backer       IFDatabase // We can have backing databases.  For now this will be nil
	persist      IFDatabase // We do need LevelDB or Bolt.  It would go here.
	doNotPersist map[string][]byte
	doNotCache   map[string][]byte
}

var _ IFDatabase = (*FDatabase)(nil)

/*************************************
 *       Stubs
 *************************************/

func (FDatabase) Close() {}

func (FDatabase) GetAddress() (IAddress, error) {
	return nil, nil
}

func (FDatabase) GetHash() IHash {
	return nil
}

func (FDatabase) GetDBHash() IHash {
	return nil
}

func (FDatabase) NumberOfSignatures() int {
	return 1
}

func (FDatabase) CustomMarshalText() ([]byte, error) {
	return nil, nil
}

func (FDatabase) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (FDatabase) UnmarshalBinary([]byte) error {
	return nil
}

func (FDatabase) UnmarshalBinaryData([]byte) ([]byte, error) {
	return nil, nil
}

func (FDatabase) DeleteKey(bucket []byte, key []byte) {}

func (FDatabase) Get(bucket string, key IHash) IBlock {
	return nil
}

func (FDatabase) GetKey(key IDBKey) IBlock {
	return nil
}

func (FDatabase) GetKeysValues(bucket []byte) (keys [][]byte, values []IBlock) {
	return nil, nil
}

func (FDatabase) IsEqual(IBlock) []IBlock {
	return nil
}

func (FDatabase) GetRaw(bucket []byte, key []byte) IBlock {
	return nil
}

func (FDatabase) String() string {
	return ""
}

func (FDatabase) Init(a ...interface{}) {}

func (FDatabase) Put(bucket string, key IHash, value IBlock) {}

func (FDatabase) PutKey(key IDBKey, value IBlock) {}

func (FDatabase) PutRaw(bucket []byte, key []byte, value IBlock) {}

/***************************************
 *       Methods
 ***************************************/

func (FDatabase) GetNewInstance() IBlock {
	return new(FDatabase)
}

// Do not hold objects in this cache in memory.  They are too big, and there
// is no interesting reason to keep them in memory.
func (db FDatabase) DoNotCache(bucket string) {
	if db.doNotCache == nil {
		db.doNotCache = make(map[string][]byte, 5)
	}
	db.doNotCache[bucket] = []byte(bucket)
	if db.backer != nil {
		db.backer.DoNotCache(bucket)
	}
	if db.persist != nil {
		db.persist.DoNotCache(bucket)
	}
}

// Do not write to disk.  These items are small, we need fast writes, and we don't need
// the overhead of writing to disk.
func (db FDatabase) DoNotPersist(bucket string) {
	if db.doNotPersist == nil {
		db.doNotPersist = make(map[string][]byte, 5)
	}
	db.doNotPersist[bucket] = []byte(bucket)
	if db.backer != nil {
		db.backer.DoNotPersist(bucket)
	}
	if db.persist != nil {
		db.persist.DoNotPersist(bucket)
	}
}

// A Backer database allows the implementation of a least recently
// used cache to purge data from memory.
func (db *FDatabase) SetBacker(b IFDatabase) {
	db.backer = b

	//TODO We should tell our backer about our DoNotCache and DoNotPersist maps.  We
	//don't now, but all we have to do is set up the databases first.
}
func (db FDatabase) GetBacker() IFDatabase {
	return db.backer
}

// A Persist database is needed to persist writes.  This is where
// one can hook up a LevelDB or Bolt database.
func (db *FDatabase) SetPersist(p IFDatabase) {
	db.persist = p

	//TODO We should tell our backer about our DoNotCache and DoNotPersist maps.  We
	//don't now, but all we have to do is set up the databases first.

}
func (db FDatabase) GetPersist() IFDatabase {
	return db.persist
}

/*****************************************************************
 * Database Key for Key/Value Databases that don't support buckets
 *****************************************************************/

type DBKey struct {
	bucket [ADDRESS_LENGTH]byte
	key    [KEY_LIMIT]byte
}

var _ IDBKey = (*DBKey)(nil)

func (k DBKey) GetBucket() []byte {
	return k.bucket[:]
}

func (k DBKey) GetKey() []byte {
	return k.key[:]
}

func makeKey(bucket []byte, key []byte) IDBKey {

	if len(bucket) > ADDRESS_LENGTH || len(key) > KEY_LIMIT {
		panic("Key or bucket provided to IFDatabase is too long")
	}

	k := new(DBKey)
	copy(k.bucket[:], bucket)
	copy(k.key[:], key)

	return k
}