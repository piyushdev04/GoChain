package blockchain

import "fmt"

type Blockchain struct {
    Blocks     []Block
    Difficulty int
}

func NewBlockchain(difficulty int) *Blockchain {
    genesis := NewBlock(0, "Genesis Block", "0")
    return &Blockchain{
        Blocks:     []Block{genesis},
        Difficulty: difficulty,
    }
}

func (bc *Blockchain) LatestBlock() Block {
    return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) {
    lastBlock := bc.LatestBlock()
    newBlock := NewBlock(lastBlock.Index+1, data, lastBlock.Hash)
    bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
    for i := 0; i < len(bc.Blocks); i++ {
        current := bc.Blocks[i]
        recalculatedHash := current.calculateHash()

        if current.Hash != recalculatedHash {
            fmt.Printf("❌ Block %d hash mismatch! Got: %s, Expected: %s\n", i, current.Hash, recalculatedHash)
            return false
        }

        if i > 0 {
            prev := bc.Blocks[i-1]
            if current.PrevHash != prev.Hash {
                fmt.Printf("❌ Block %d prevHash mismatch!\n", i)
                return false
            }
        }
    }

    fmt.Println("✅ Blockchain is valid.")
    return true
}