package kapo

import (
	"errors"

	"github.com/boltdb/bolt"
)

var (
	HeadDoesNotExist = errors.New("HEAD does not exist")
)

const (
	storeFile         = "kapo.db"
	storeBlocksBucket = "blocks"
	storeHeadKey      = "HEAD"
)

// NOTE: Store could/should be an interface

type Store struct {
	db *bolt.DB
}

func (s *Store) Open() error {
	err := s.open(storeFile)
	if err != nil {
		return err
	}

	// Create buckets if do not exist
	err = s.ensureBuckets()
	if err != nil {
		return err
	}

	// Create genesis block if does not exist
	err = s.ensureGenesisBlock()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Close() error {
	if s.db != nil {
		return s.db.Close()
	}

	return nil
}

func (s *Store) open(filename string) error {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) ensureBuckets() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storeBlocksBucket))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Store) ensureGenesisBlock() error {
	_, err := s.Head()
	if err != nil {
		if err == HeadDoesNotExist {
			s.PutBlock(DefaultGenesisBlock())
		} else {
			return err
		}
	}

	return nil
}

func (s *Store) Head() (Hash, error) {
	var id []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storeBlocksBucket))
		id = b.Get([]byte(storeHeadKey))
		return nil
	})
	if err != nil {
		return Hash{}, nil
	}

	if len(id) == 0 {
		return Hash{}, HeadDoesNotExist
	}

	return ToHash(id), nil
}

func (s *Store) GetBlock(id Hash) (*Block, error) {
	block := &Block{}

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storeBlocksBucket))
		encodedBlock := b.Get(id.Bytes())
		err := block.Deserialize(encodedBlock)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *Store) PutBlock(block *Block) error {
	id := block.Hash.Bytes()

	data, err := block.Serialize()
	if err != nil {
		return err
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storeBlocksBucket))

		err = b.Put(id, data)
		if err != nil {
			return err
		}

		err = b.Put([]byte(storeHeadKey), id)
		if err != nil {
			return err
		}

		return nil
	})
}
