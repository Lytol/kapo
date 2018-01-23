package kapo

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock}}
}

func (bc *Blockchain) AddBlock(txs []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(txs, prevBlock.ID)
	bc.Blocks = append(bc.Blocks, newBlock)
}
