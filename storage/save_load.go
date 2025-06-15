package storage

import (
    "encoding/json"
    "fmt"
    "gochain/blockchain"
    "io/ioutil"
    "os"
)

const FileName = "blockchain.json"

func SaveToDisk(bc *blockchain.Blockchain) error {
    data, err := json.MarshalIndent(bc.Blocks, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(FileName, data, 0644)
}

func LoadFromDisk(difficulty int) *blockchain.Blockchain {
    if _, err := os.Stat(FileName); os.IsNotExist(err) {
        fmt.Println("🆕 No blockchain found — creating genesis block.")
        return blockchain.NewBlockchain(difficulty)
    }

    data, err := ioutil.ReadFile(FileName)
    if err != nil {
        panic(err)
    }

    var blocks []blockchain.Block
    err = json.Unmarshal(data, &blocks)
    if err != nil {
        panic(err)
    }

    fmt.Println("📂 Blockchain loaded from disk.")
    return &blockchain.Blockchain{
        Blocks:     blocks,
        Difficulty: difficulty,
    }
}
