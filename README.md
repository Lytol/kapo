Kapo
====

Just messing around with blockchains and distributed ledgers using Go


Goals
-----

- Small and minimal, do less
- Single binary with no runtime dependencies
- Permissioned writes / public reads
- Smart contracts
- Storage oriented (documents, etc)


Architecture
------------

- [BoltDB](https://github.com/boltdb/bolt) for the ledger database


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

- [ ] Refactor genesis block into something sane
- [ ] Addresses
- [ ] Scripting
- [ ] Networking
- [ ] Permissioned blockchain (read all / permissioned write)
- [ ] Proof of work is probably not the right consensus mechanism. Research proof of stake and other consensus algo's that incentivize storing and verifying transactions and documents.
- [ ] Merkle tree
