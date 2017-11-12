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

这里提到哈希值，我们开始可以不需要关注到比较复杂的PBFT/POW等共识算法，可以仅仅将几个值关联在一起，计算SHA-256。像这样

    headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

#### 区块链
区块链是带有特定数据结构的数据库：它是一个有序的、反向链接的列表（back-linked list）这意味着区块是按照插入的顺利存储的，并且每个区块都链接到前面一个区块。这种结构允许快速地获取最新的区块，也可以非常高效地通过哈希值获得其对应的区块。

#### POW
关于哈希值的计算，比特币采用HASHCASH算法，是一种在早起用语垃圾邮件过滤的算法。，简单的暴力求解，通过。通过检验，递增计数器，确认要求，反馈。

在Hashcash算法的原始实现当中，对哈希值的要求是“前20位必须为0”。在比特币当中，要求随时间变化有所调整，因为，根据设计，必须每10分钟产生一个区块，不管算力如何增加并有越来越多的矿工加入网络当中。

在比特币当中，“目标位数”（targetBits）是存储在区块头部数据用以指示挖矿难度的指标。目前，我们并不打算实现难度可调的算法。因此，我们可以将难度系数定义为一个全局常量。

设置一个 ProofOfWork 的结构体：

    type ProofOfWork struct {
        block  *Block
        target *big.Int
    }

之所以用Big整型来定义“target”在于我们将哈希值与目标比较的方式。我们将一个哈希值转换为一个Big整型然后检验其是否小于目标值。
可以将target理解成一个范围的上边界：假如一个数（一个哈希值）比这个边界小，有效，反之，则无效。减小边界的数值，会导致更少的有效数字的个数，这样得到一个有效哈希值的难度将加大。


Refference:  [blockchain_go](https://github.com/Jeiwan/blockchain_go)