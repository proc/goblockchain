package main

import (
	"crypto/sha256"
	"crypto/rand"
	"math"
	"math/big"
	"fmt"
)

type Block struct {
	PreviousHash []byte
	NextHash string
	SequenceNumber int
	Hash []byte
	Data []byte
}

func (b *Block) GenerateHash() []byte {
	if b.Hash != nil {
		fmt.Println("Hash already exists")
		return nil
	}

	h := sha256.New()
	nonce, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	data := fmt.Sprintf("%s%s", b.Header(), nonce.String())
	fmt.Println("Generating hash with data: ", data)
	h.Write([]byte(data))
	b.Hash = h.Sum(nil)

	return b.Hash
}

func (b Block) Header() string {
	// fmt.Printf("previous hash %s\n", b.PreviousHash)
	return fmt.Sprintf("%s%x", b.Data, b.PreviousHash)
}

func (b Block) HashString() string {
	return fmt.Sprintf("%x", b.Hash)
}

func NewBlock(data []byte) *Block {
	block := &Block{
		Data: data,
	}
	return block
}

type BlockChain struct {
	Genesis *Block
	Last *Block
	Blocks map[string]*Block
}

func NewBlockChain() *BlockChain {
	genesis := &Block{
		Data: []byte("genesis block"),
	}
	genesis.GenerateHash()
	return &BlockChain{
		Genesis: genesis,
		Last: genesis,
		Blocks: map[string]*Block{
			fmt.Sprintf("%x", genesis.Hash): genesis,
		},
	}
}
func (bc *BlockChain) AddBlock(b *Block) error {
	b.PreviousHash = bc.Last.Hash
	b.SequenceNumber = len(bc.Blocks)
	b.GenerateHash()
	fmt.Printf("Adding Block Number: %d\n", b.SequenceNumber)
	bc.Last = b
	bc.Blocks[b.HashString()] = b

	return nil
}

func (bc BlockChain) Print() {
	fmt.Println("-----------------------------------")
	bc.PrintRecursive(bc.Last)
	fmt.Println("-----------------------------------")
}

func (bc BlockChain) PrintRecursive(b *Block) {
	if b.HashString() == bc.Genesis.HashString() {
		fmt.Printf("[%x](Genesis)\n", b.Hash)
		return
	}
	fmt.Printf("[%x]\n|\n", b.Hash)
	bc.PrintRecursive(bc.Blocks[fmt.Sprintf("%x", b.PreviousHash)])
}

func main() {
	blockChain := NewBlockChain()
	fmt.Printf("Genesis block created: %x\n", blockChain.Genesis.Hash)

	err := blockChain.AddBlock(NewBlock([]byte("just some random data")))
	if err != nil {
		panic(err)
	}
	err = blockChain.AddBlock(NewBlock([]byte("just another block")))
	if err != nil {
		panic(err)
	}
	err = blockChain.AddBlock(NewBlock([]byte("2 + 2 = 4")))
	if err != nil {
		panic(err)
	}
	err = blockChain.AddBlock(NewBlock([]byte("snack snack")))
	if err != nil {
		panic(err)
	}
	err = blockChain.AddBlock(NewBlock([]byte("esper")))
	if err != nil {
		panic(err)
	}
	err = blockChain.AddBlock(NewBlock([]byte("bing")))
	if err != nil {
		panic(err)
	}
	blockChain.Print()
}
