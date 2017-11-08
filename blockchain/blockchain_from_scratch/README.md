Blockchain from scratch
---

## Introduction

通过底层来开发区块链的原型（BitCoin的），加入地理信息作为供应链或者物联网方面的验证。

## 区块链的必须模块

#### 区块

区块是区块链的核心，类似于一个数据集，包含有时间，哈希值，上一个区块信息。。。类似于

    type Block struct {
        Timestamp     int64
        Transactions  []*Transaction
        PrevBlockHash []byte
        Hash          []byte
        Nonce         int
        Height        int
    }

这样,这里提到哈希值.



Refference:  [blockchain_go](https://github.com/Jeiwan/blockchain_go)