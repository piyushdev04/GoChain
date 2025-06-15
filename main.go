package main

import (
    "bufio"
    "fmt"
    "gochain/blockchain"
    "gochain/storage"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    bc := storage.LoadFromDisk(2) // Difficulty = 2

    for {
        fmt.Println("\n📘 Welcome to GoChain ⛓️")
        fmt.Println("[1] View Blockchain")
        fmt.Println("[2] Add Block")
        fmt.Println("[3] Tamper Block (for testing)")
        fmt.Println("[4] Validate Blockchain")
        fmt.Println("[5] Exit")
        fmt.Print("Choose an option: ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "1":
            printBlockchain(bc)
        case "2":
            fmt.Print("Enter block data: ")
            data, _ := reader.ReadString('\n')
            data = strings.TrimSpace(data)
            bc.AddBlock(data)
            storage.SaveToDisk(bc)
        case "3":
            fmt.Print("Enter index to tamper: ")
            idxStr, _ := reader.ReadString('\n')
            idxStr = strings.TrimSpace(idxStr)
            idx, err := strconv.Atoi(idxStr)
            if err == nil && idx >= 0 && idx < len(bc.Blocks) {
                bc.Blocks[idx].Data = "[tampered]"
                fmt.Println("🔨 Block tampered!")
                storage.SaveToDisk(bc)
            } else {
                fmt.Println("❌ Invalid index")
            }
        case "4":
            bc.IsValid()
        case "5":
            fmt.Println("👋 Exiting GoChain.")
            return
        default:
            fmt.Println("❌ Invalid choice.")
        }
    }
}

func printBlockchain(bc *blockchain.Blockchain) {
    fmt.Println("📦 Blockchain Contents:\n")
    for _, block := range bc.Blocks {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Println("---------------------------")
    }
}
