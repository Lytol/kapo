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
- [ ] Refactor crypto into its own package
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


Sample Account
--------------

DO NOT USE THIS FOR ANYTHING OTHER THAN TESTING

Address:    d2e65ab43e6fb51852d261f90ae3c62455461e9055b4a1c6bea965190c4770ec
PrivateKey: ed0f33026b50ab641ade1ac428bfded42b9047c54d5d420d1efc50d7793175f4
PublicKey:  118f3ddbe11e271af3ce5e9fdfb1c664f5db82752c44c1571aab8036626d57e99b6018f167a77ac9717b7b5b13d7645932bcd75e75e33f70d15b87912c81528d
