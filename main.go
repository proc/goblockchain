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
	data := fmt.Sprintf("%s%s", b.Data, nonce.String())
	fmt.Println("Generating hash with data: ", data)
	h.Write([]byte(data))
	b.Hash = h.Sum(nil)

	return b.Hash
}

func (b Block) HashString() string {
	return fmt.Sprintf("%x", b.Hash)
}

func NewBlock(data []byte) *Block {
	block := &Block{
		Data: data,
	}
	hash := block.GenerateHash()
	fmt.Println("New Block created: ", fmt.Sprintf("%x", hash))
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
	fmt.Printf("Adding Block: %x\n", b.Hash)
	b.PreviousHash = bc.Last.Hash
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

	blockChain.AddBlock(NewBlock([]byte("just some random data")))
	blockChain.AddBlock(NewBlock([]byte("just another block")))
	blockChain.AddBlock(NewBlock([]byte("2 + 2 = 4")))
	blockChain.AddBlock(NewBlock([]byte("snack snack")))
	blockChain.AddBlock(NewBlock([]byte("esper")))
	blockChain.AddBlock(NewBlock([]byte("bing")))

	blockChain.Print()
}