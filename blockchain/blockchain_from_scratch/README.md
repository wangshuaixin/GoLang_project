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

#### 存储

在blockchain结构中放入一个数据库链接。因为我们希望打开一次以后只要程序在运行就一直保持打开。并且之后可以直接通过写blockchain的方法来访问数据库


#### 数据结构

简单地讲，比特币内核用了两个“Buckets”来存储数据：

+ blocks：存储区块链中描述区块的元数据
+ chainstate：存储链的状态，包括所有未成交的交易输出（Outputs）记录和一些元数据

并且，每一个blocks在硬盘中以单独的文件来保存，这样做是为了提高性能：读取一个单独的区块并不需要载入所有（或者部分）区块到内存。

不过我们并不这样实现。在blocks 当中，Key -> Value 数据对如下：

    'b' + 32-byte block hash -> block index record
    'f' + 4-byte file number -> file information record
    'l' -> 4-byte file number: the last block file number used
    'R' -> 1-byte boolean: whether we're in the process of reindexing
    'F' + 1-byte flag name length + flag name string -> 1 byte boolean: various flags that can be on or off
    't' + 32-byte transaction hash -> transaction index record在chainstate 
    当中，Key -> Value 数据对如下：
    'c' + 32-byte transaction hash -> unspent transaction output record for that transaction
    'B' -> 32-byte block hash: the block hash up to which the database represents the unspent transaction outputs

因为目前我们还没有交易记录，我们的数据库当中仅有 blocks 这个bucket。

并且，正如前面所提到的，我们将全部的数据库存储在一个文件当中，并不将单个区块存储在单个文件当中。因此，我们也没有什么信息是与文件数（file number）有个的，所能够利用的 key -> value 数据对如下：

    'b' + 32-byte block-hash -> Block structure (serialized)区块结构体（序列化后的）
    'l' -> the hash of the last block in a chain（区块链中最近一个区块的哈希值）
    以上是目前开始实现数据存储机制所需要了解的全部内容。

#### 交易记录

以比特币为例，交易记录的主要特点如下：

    没有账户
    没有资产负债表
    没有地址
    没有货币
    没有支付方和接收方

一条交易记录中是所有输入和输出的合并：

    type Transaction struct {
        ID   []byte
        Vin  []TXInput
        Vout []TXOutput
    }

但是也要注意下面几点：

    有一些输出并不与输入相对应
    在一个交易记录当中，输入可以与来自不同交易记录的输出相对应
    一个输入只能对于一个输出

交易记录的输出的结构体里面包括两个内容，一个是币值，一个是一段脚本语言。在比特币内部使用一个叫做Script的脚本语言，用来定义输出的锁定和解锁逻辑。非常接近底层，这样设计是为了避免滥用和攻击。

交易记录的输入包括一个交易ID，还有币值和脚本语言。ScriptSig 字段是一段向输出的 ScriptPubKey 字段中提供数据的脚本。如果数据正确的话，输出将会被解锁，然后其含的价值（value）可以用来产生新的输出；如果不正确的话，输出将无法被输入引用，无法建立连接。这个机制是为了保证用户不能花属于别人的比特币。

所以，输出是“币”存储的地方。每一个输出会带一段解锁的脚本（字符串），决定了解锁输出的逻辑。每一个新的交易记录只是有一个输入和一个输出。一个输入对应一个从前一个交易记录（transaction）来的输出并提供用于解锁（unlocking）输出的数据以便使用输出中的币值（value）来创建新的输出（outputs）

由于输出和输出的类似性，所以添加一个coinbase交易，当一个矿工开始挖坑（mining a block），将在区块当中添加币基交易记录（coinbase transaction）。币基交易记录是一种特殊的交易记录，并不需要已经存在的输出就能够产生。它可以从无直接创造出输出。相当于不需要鸡的蛋。这是矿工在挖新区块时获得的奖励。


#### Merkle Tree
比特币用一个更加复杂的技术：它采用Merkle tree来组织一个区块中的所有交易记录，然后用树的根哈希值来确保PoW系统的运行。这种方法能够让我们快速的检验一个区块是否包含确定的交易记录，只要有根哈希值而不需要下载所有的交易记录。

关于merkle tree可以看这一篇博客

https://www.cnblogs.com/fengzhiwu/p/5524324.html
 

#### UTXO
这里指没有花费交易记录的输出，当然，当我们检查余额时，我们并不需要全部，只需要那些我们有私钥可以解锁的部分。通过TX结构体中的ScriptSig和ScriptPubKey等等进行比较。

关于寻找没有花费的交易记录：


    func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
        var unspentTXs []Transaction
        //一个未花费交易记录
        spentTXOs := make(map[string][]int)
        bci := bc.Iterator()
        //查看blockchain
        for {
            block := bci.Next()
            //查找下一个
            for _, tx := range block.Transactions {
            txID := hex.EncodeToString(tx.ID)
            //获取ID
            Outputs:
            for outIdx, out := range tx.Vout {
                // 由于交易记录存在在区块当中，我们必须检查区块链中的每一个区块
                if spentTXOs[txID] != nil {
                for _, spentOut := range spentTXOs[txID] {
                    //当一个输出由我们用来选择未花费交易记录的地址上的锁，那么这个输出就是我们想要的。
                    //但是在我们取得它之前我们需要确认它是否已经与一个输入相关联
                    if spentOut == outIdx {
                    continue Outputs
                    }
                }
                }

                if out.CanBeUnlockedWith(address) {
                unspentTXs = append(unspentTXs, *tx)
                }
            }
            //我们将忽略哪些已经与输入相关的的输出
            //它们的价值已经转移到其它的输出，所以不能再统计它们
            //检查输出以后，我们手机所有那些可以结果被提供的地址锁上的输出的输入，这对币基交易记录不适用，因为它们不解锁任何输出
            if tx.IsCoinbase() == false {
                for _, in := range tx.Vin {
                if in.CanUnlockOutputWith(address) {
                    inTxID := hex.EncodeToString(in.Txid)
                    spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
                }
                }
            }
            }

            if len(block.PrevBlockHash) == 0 {
            break
            }
        }

    return unspentTXs
    }

下面的函数返回一系列包含未花费输出的交易记录。为了计算余额，我们需要额外的一个以交易记录为参数并只返回输出的函数

    func (bc *Blockchain) FindUTXO(address string) []TXOutput {
        var UTXOs []TXOutput
        unspentTransactions := bc.FindUnspentTransactions(address)

        for _, tx := range unspentTransactions {
                for _, out := range tx.Vout {
                        if out.CanBeUnlockedWith(address) {
                                UTXOs = append(UTXOs, out)
                        }
                }
        }

        return UTXOs
    }

#### send coin
就是创建一个UNTXOTransaction，在创建新的输出之前，我们首先找出所有的未消费输出并且确保它们存有足够的币值。这是FindSpendableOutputs 方法的功能

然后需要两个输出

一个用接受者的地址进行锁定。这是实际需要转移到其它地址的币值。

一个用发送者的地址进行锁定。这是找零。当且仅当剩余未花费输出持有的总币值比新的交易记录所要求的多。记住：输出是不可分割的。

    func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {}

这个方法遍历所有的未花费交易记录并累计它们的币值。当累计的币值大于或者等于我们所有转移的量时，它停止工作然后返回累计币值（accumulated value）以及按交易记录ID进行分组的输出索引。我们不打算取比我们打算要花费的多。


- - -
Refference: 

https://github.com/Jeiwan/blockchain_go

https://github.com/bitcoin/bitcoin

https://bitcoincore.org/bitcoin.pdf

https://en.wikipedia.org/wiki/Merkle_tree
- - -