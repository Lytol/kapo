Kapo
====

Just messing around with blockchains and distributed ledgers using Go

Usage
-----

- add : add a transaction to the blockchain (requires private key, account, script, optional data)

Goals
-----

- Small and minimal, do less
- Single binary with no runtime dependencies
- Smart contracts
- Storage oriented (documents, etc)


Architecture
------------

- [BoltDB](https://github.com/boltdb/bolt) for the ledger database (transactions, blocks)


Resources
---------

- [Bitcoin: A Peer-to-Peer Electronic Cash System](https://bitcoin.org/bitcoin.pdf)
- [Ethereum White Paper](https://github.com/ethereum/wiki/wiki/White-Paper)
- [Ethereum Source Code](https://github.com/ethereum/go-ethereum/)
- [Storj: A Peer-to-Peer Cloud Storage Network](https://storj.io/storj.pdf)
- [IPFS](https://ipfs.io/)
- [Ivan Kuznetsov - Building Blockchain in Go](https://jeiwan.cc/tags/blockchain/)
- [Kademlia: A Peer-to-peer Information System Based on the XOR Metric](https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf)


TODO
----

- [ ] Addresses
- [ ] Refactor mining into "seal" per Ethereum
- [ ] Scripting
- [ ] Networking
- [ ] Permissioned blockchain (read all / permissioned write)
- [ ] Proof of work is probably not the right consensus mechanism. Research proof of stake and other consensus algo's that incentivize storing and verifying transactions and documents.
- [ ] Merkle tree


Notes
-----

Transactions:

  Creation:

    - Signature = SIGN(Hash(Address + Script + Data)) + PubKey
    - Hash = Hash(Address + Script + Data + Signature)

  Verification:
    - Ensure Hash
    - Ensure signature from public key
    - Ensure account from public key

Blocks:

  Creation
    - Hash = Hash(Tx1.Hash + Tx2.Hash + ... + TxN.Hash)
    - Consensus: PoW, etc

  Verification
    - Ensure Hash
    - Ensure PoW
