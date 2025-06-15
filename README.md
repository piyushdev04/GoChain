# GoChain — A Simple Blockchain in Go

GoChain is a simple Blockchain. It uses Proof-of-Work to mine blocks, stores data in a JSON file, and includes an interactive CLI to visualize how blockchains work at a fundamental level.

---

## Features

- Custom Blockchain structure with chaining of blocks
- Proof-of-Work mining (adjustable difficulty)
- Tamper detection with hash integrity validation
- Persistent storage via `blockchain.json`
- Interactive CLI for adding/viewing/tampering blocks
- Built-in testing for detecting tampering or corruption

---

## Technical Concepts
1. Block Structure
- Each block contains:

- Index: Position in chain

- Timestamp: Creation time

- Data: Payload (e.g., transaction)

- PrevHash: Hash of previous block

- Hash: Current block hash

- Nonce: PoW nonce used to mine this block

2. Proof-of-Work
A simple PoW algorithm ensures each block’s hash starts with a fixed number of 0s based on difficulty. Mining finds a nonce that satisfies this.

3. Tampering Detection
Modifying any part of a block (like Data) changes its hash. The blockchain validator recalculates and compares all hashes to ensure integrity.