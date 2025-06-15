package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Hash      string
    Nonce     int
}

func NewBlock(index int, data string, prevHash string) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().Format(time.RFC3339),
        Data:      data,
        PrevHash:  prevHash,
        Nonce:     0,
    }
    block.MineBlock(2) // Default difficulty
    return block
}

func (b *Block) calculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

func (b *Block) MineBlock(difficulty int) {
    prefix := ""
    for i := 0; i < difficulty; i++ {
        prefix += "0"
    }

    for {
        b.Hash = b.calculateHash()
        if b.Hash[:difficulty] == prefix {
            fmt.Printf("⛏️  Mined Block %d with hash %s\n", b.Index, b.Hash)
            break
        }
        b.Nonce++
    }
}
