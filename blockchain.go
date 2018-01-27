package kapo

type Blockchain struct {
	store *Store
}

func NewBlockchain() (*Blockchain, error) {
	store := &Store{}

	err := store.Open()
	if err != nil {
		return nil, err
	}

	return &Blockchain{store}, nil
}

func (bc *Blockchain) AddBlock(txs []*Transaction) error {
	head, err := bc.store.Head()
	if err != nil {
		return err
	}

	block := NewBlock(txs, head)

	return bc.store.PutBlock(block)
}

func (bc *Blockchain) Head() (Hash, error) {
	return bc.store.Head()
}

type BlockchainIterator struct {
	blockchain  *Blockchain
	currentHash Hash
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	head, err := bc.store.Head()
	if err != nil {
		panic(err) // TODO: Don't panic
	}

	return &BlockchainIterator{bc, head}
}

func (i *BlockchainIterator) Next() *Block {
	if i.currentHash == DefaultGenesisBlock().Previous {
		return nil
	}

	block, err := i.blockchain.store.GetBlock(i.currentHash)
	if err != nil {
		panic(err) // TODO: Don't panic
	}

	i.currentHash = block.Previous

	return block
}
