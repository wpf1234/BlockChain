package blockchain

import (
	"BlockChain/block"
	"BlockChain/pow"
	"BlockChain/transaction"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"
)

// 实现一个区块链

//// 声明一个表示区块链的结构体
////type BlockChain struct {
////	blocks []*block.Block   // blocks 为保存区块的切片
////}
//
//type BlockChain struct {
//	tip []byte
//	db  *bolt.DB
//}
//
///*
//   简单的区块链实现如下
//*/
//
//// 加入区块前先创建一个创世区块
////func GenesisBlock() *block.Block{
////	return block.NewBlock("Genesis Block",[]byte{})
////}
//
//// 把区块添加到区块链中
////func (bc *BlockChain) AddBlock(data string){
////	prevBlock:=bc.blocks[len(bc.blocks)-1]    // 前一个区块
////	newBlock:=block.NewBlock(data,prevBlock.Hash)
////	bc.blocks=append(bc.blocks,newBlock)
////}
//
///*
//   POW 算法，创建创世区块，把区块添加到区块链，方法如下
//*/
//
//const dbFile = "blockChain.db"
//const blocksBucket = "blocks"
////const genesisCoinbaseData = "Time 2020/03/05,learn block"
//
////func GenesisBlock() *block.Block {
////	return pow.NewBlock("创世区块: ", []byte{})
////}
//
//// 添加链上交易
//func NewGenesisBlock(coinbase *transaction.Transaction) *block.Block {
//	return pow.NewBlock([]*transaction.Transaction{coinbase}, []byte{})
//}
//
////func (bc *BlockChain) AddBlock(data string) {
////	prevBlock := bc.blocks[len(bc.blocks)-1] // 前一个区块
////	newBlock := pow.NewBlock(data, prevBlock.Hash)
////	bc.blocks = append(bc.blocks, newBlock)
////}
//
////func (bc *BlockChain) AddBlock(data string) {
////	var lastHash []byte
////	// 只读的方式浏览数据库，获取当前区块链顶端区块的哈希，为加入下一区块做准备
////	err := bc.db.View(func(tx *bolt.Tx) error {
////		b := tx.Bucket([]byte(blocksBucket))
////		// 通过 "l" 键获取区块链顶端区块哈希
////		lastHash = b.Get([]byte("l"))
////
////		return nil
////	})
////
////	if err != nil {
////		log.Error(err)
////		return
////	}
////
////	// 求出新的区块
////	newBlock := pow.NewBlock(data, lastHash)
////	// 把新的区块加入区块链中
////	err = bc.db.Update(func(tx *bolt.Tx) error {
////		b := tx.Bucket([]byte(blocksBucket))
////		err := b.Put(newBlock.Hash, newBlock.Serialize())
////		if err != nil {
////			log.Error("新的区块加入失败: ", err)
////			return nil
////		}
////		err = b.Put([]byte("l"), newBlock.Hash)
////		bc.tip = newBlock.Hash
////
////		return nil
////	})
////}
//
//// 添加交易
//func (bc *BlockChain) AddBlock(transaction []*transaction.Transaction) {
//	var lastHash []byte
//	// 只读的方式浏览数据库，获取当前区块链顶端区块的哈希，为加入下一区块做准备
//	err := bc.db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		// 通过 "l" 键获取区块链顶端区块哈希
//		lastHash = b.Get([]byte("l"))
//
//		return nil
//	})
//
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	// 求出新的区块
//	newBlock := pow.NewBlock(transaction, lastHash)
//	// 把新的区块加入区块链中
//	err = bc.db.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		err := b.Put(newBlock.Hash, newBlock.Serialize())
//		if err != nil {
//			log.Error("新的区块加入失败: ", err)
//			return nil
//		}
//		err = b.Put([]byte("l"), newBlock.Hash)
//		bc.tip = newBlock.Hash
//
//		return nil
//	})
//}
//
//// 初始化区块链，默认存储了创世区块
//
////func NewBlockChain() *BlockChain {
////	//return &BlockChain{[]*block.Block{GenesisBlock()}}
////
////	/*
////	  1.打开数据文件
////	  2.检查文件里面是否已经存储了一个区块链
////	  3.如果有，创建一个新的 BlockChain 实例，设置 BlockChain 实例的 tip 为数据库中存储的最后一个块的哈希值
////	  4.如果没有，创建创世块，存储到数据库中，将创世块哈希保存为最后一个快的哈希，创建一个新的 BlockChain 实例
////	    其 tip 指向创世块(tip 存储的是最后一块的哈希)
////	*/
////
////	var tip []byte
////	// 打开一个数据库文件，如果文件不存在则创建该文件
////	db, err := bolt.Open(dbFile, 0600, nil)
////	if err != nil {
////		log.Error("打开数据库文件失败: ", err)
////		return nil
////	}
////	// 读写操作数据库
////	err = db.Update(func(tx *bolt.Tx) error {
////		b := tx.Bucket([]byte(blocksBucket))
////		// 查看名字为 blocksBucket 的 Bucket 是否存在
////		if b == nil {
////			// 不存在，从头创建
////			//fmt.Println("不存在的区块链，需要重新创建一个区块链......")
////			//coinbaseData:="创建新的一个区块链..."
////			genesis := GenesisBlock()                       // 创建创世区块
////			b, err := tx.CreateBucket([]byte(blocksBucket)) // 创建名为 blockBucket 的桶
////			if err != nil {
////				log.Error("创建新的 bucket 失败: ", err)
////				return nil
////			}
////			// 写入键值对，区块哈希对应序列化后的区块
////			err = b.Put(genesis.Hash, genesis.Serialize())
////			if err != nil {
////				log.Error("写入数据失败: ", err)
////				return nil
////			}
////			// 指向最后一块区块,这里也就是创世区块
////			tip = genesis.Hash
////		} else {
////			// 存在
////			// 通过键“l” 映射出顶端区块的 hash
////			tip = b.Get([]byte("l"))
////		}
////		return nil
////	})
////
////	// 此时 Blockchain 结构体字段已经变成这样了
////	bc := BlockChain{
////		tip: tip,
////		db:  db,
////	}
////	return &bc
////}
//// 添加交易
//func NewBlockChain(address string) *BlockChain {
//	//return &BlockChain{[]*block.Block{GenesisBlock()}}
//
//	/*
//	  1.打开数据文件
//	  2.检查文件里面是否已经存储了一个区块链
//	  3.如果有，创建一个新的 BlockChain 实例，设置 BlockChain 实例的 tip 为数据库中存储的最后一个块的哈希值
//	  4.如果没有，创建创世块，存储到数据库中，将创世块哈希保存为最后一个快的哈希，创建一个新的 BlockChain 实例
//	    其 tip 指向创世块(tip 存储的是最后一块的哈希)
//	*/
//
//	var tip []byte
//	// 打开一个数据库文件，如果文件不存在则创建该文件
//	db, err := bolt.Open(dbFile, 0600, nil)
//	if err != nil {
//		log.Error("打开数据库文件失败: ", err)
//		return nil
//	}
//	// 读写操作数据库
//	err = db.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		// 查看名字为 blocksBucket 的 Bucket 是否存在
//		if b == nil {
//			// 不存在，从头创建
//			fmt.Println("不存在的区块链，需要重新创建一个区块链......")
//			coinbaseData := "创建新的一个区块链..."
//			//genesis := GenesisBlock()                       // 创建创世区块
//			// 此时的创世 区块就要包含交易 coinbaseTX
//			cbtx := transaction.NewCoinbaseTX(address, coinbaseData)
//			genesis := NewGenesisBlock(cbtx)
//			b, err := tx.CreateBucket([]byte(blocksBucket)) // 创建名为 blockBucket 的桶
//			if err != nil {
//				log.Error("创建新的 bucket 失败: ", err)
//				return nil
//			}
//			// 写入键值对，区块哈希对应序列化后的区块
//			err = b.Put(genesis.Hash, genesis.Serialize())
//			if err != nil {
//				log.Error("写入数据失败: ", err)
//				return nil
//			}
//			// "l"键对应区块链顶端区块的哈希值
//			err = b.Put([]byte("l"), genesis.Hash)
//			if err != nil {
//				log.Error(err)
//				return nil
//			}
//			// 指向最后一块区块,这里也就是创世区块
//			tip = genesis.Hash
//		} else {
//			// 存在
//			// 通过键“l” 映射出顶端区块的 hash
//			tip = b.Get([]byte("l"))
//		}
//		return nil
//	})
//
//	// 此时 Blockchain 结构体字段已经变成这样了
//	bc := BlockChain{
//		tip: tip,
//		db:  db,
//	}
//	return &bc
//}
//
//// 工厂模式
////func (bc *BlockChain) Blocks() []*block.Block {
////	return bc.blocks
////}
//func (bc *BlockChain) DB() *bolt.DB {
//	return bc.db
//}
//
//// 把区块添加进区块链，挖矿
//func (bc *BlockChain) MineBlock(transactions []*transaction.Transaction){
//	var lastHash []byte
//	//只读的方式浏览数据库，获取当前区块链顶端区块的哈希，为加入下一区块做准备
//	err := bc.db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		lastHash = b.Get([]byte("l"))	//通过键"l"拿到区块链顶端区块哈希
//
//		return nil
//	})
//	if err != nil {
//		log.Error(err)
//	}
//
//	//prevBlock := bc.Blocks[len(bc.Blocks)-1]
//	//求出新区块
//	newBlock := pow.NewBlock(transactions,lastHash)
//	// bc.Blocks = append(bc.Blocks,newBlock)
//	//把新区块加入到数据库区块链中
//	err = bc.db.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		err := b.Put(newBlock.Hash,newBlock.Serialize())
//		if err != nil {
//			log.Error(err)
//		}
//		err = b.Put([]byte("l"),newBlock.Hash)
//		bc.tip = newBlock.Hash
//
//		return nil
//	})
//
//}
//
////----------------------------------------------------------------------------------------
//
//// 迭代器
//type BlockchainIterator struct {
//	currentHash []byte
//	db          *bolt.DB
//}
//
//// 当需要便利当前区块链时，创建一个次区块链的迭代器
//func (bc *BlockChain) Iterator() *BlockchainIterator {
//	bci := &BlockchainIterator{
//		currentHash: bc.tip,
//		db:          bc.db,
//	}
//	return bci
//}
//
//// 迭代器的任务就是返回链中下一个区块
//func (i *BlockchainIterator) Next() *block.Block {
//	var Block *block.Block
//
//	// 只读方式打开区块链数据库
//	err := i.db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		// 获取数据库中当前区块哈希对应的被序列化的区块
//		encodeBlock := b.Get(i.currentHash)
//		// 反序列化
//		Block = block.DeserializeBlock(encodeBlock)
//
//		return nil
//	})
//	if err != nil {
//		log.Error("获取区块并反序列化失败: ", err)
//		return nil
//	}
//
//	// 把迭代器中的当前区块哈希设置为上一区块的哈希，实现迭代作用
//	i.currentHash = Block.PrevBlockHash
//
//	return Block
//}
//
//// 在区块链上找到每一个区块中属于 address 用户的未花费交易输出，返回未花费输出的交易切片
//func (bc *BlockChain) FindUnspentTransaction(address string) []transaction.Transaction {
//	var unspent []transaction.Transaction
//	// 创建一个 map，存储已经花费了的交易输出
//	spent := make(map[string][]int)
//	// 因为要在链上遍历区块，所以要使用到迭代器
//	bci := bc.Iterator()
//	for {
//		block := bci.Next() // 迭代
//		for _, tx := range block.Transactions {
//			// 把交易 ID 转换成 string，方便存入 map 中
//			txID := hex.EncodeToString(tx.ID)
//
//			// 标签
//		Outputs:
//			// 遍历当前交易中的输出切片，取出交易输出
//			for outIdx, out := range tx.Vout {
//				// 在已经花费了的交易输出 map 中，如果没有找到对应的交易输出，则表示当前交易的输出未花费
//				// 反之如下
//				if spent[txID] != nil {
//					// 存在当前交易的输出中有已经花费的交易输出
//					// 则我们遍历 map 中保存的该交易 ID 对应的输出的 index
//					// 提示: 这里的已经花费的交易输出 index 实际就是输入 TXInput 结构体中的 Vout 字段
//					for _, spentOut := range spent[txID] {
//						// 首先要清楚当前交易输出是一个切片，里面有很多输出
//						// 如果 map 里存储的引用的输出和我们当前遍历到的输出 index 重合，则表示该输出被引用了
//						if spentOut == outIdx {
//							// 相同，我们就继续遍历下一轮，找到未被引用的输出
//							continue Outputs
//						}
//					}
//				}
//
//				// 到这里是得到次交易输出切片中未被引用的输出
//				// 这里就要从这些未被引用的输出中筛选属于该用户 address 地址的输出
//				if out.CanBeUnlockedWith(address) {
//					unspent = append(unspent, *tx)
//				}
//			}
//
//			// 判断是否为 coinbase 交易
//			if tx.IsCoinbase() == false {
//				// 如果不是，则遍历当前交易的输入
//				for _, in := range tx.Vin {
//					// 如果当前交易的输入是被该地址 address 所花费的
//					// 则在 map 上记录该输入引用的该地址对应的交易输出
//					if in.CanUnlockOutputWith(address) {
//						inTxID := hex.EncodeToString(in.Txid)
//						spent[inTxID] = append(spent[inTxID], in.Vout)
//					}
//				}
//			}
//		}
//
//		// 退出 for 循环的条件就是遍历到创世区块后
//		if len(block.PrevBlockHash) == 0 {
//			break
//		}
//	}
//	return unspent
//}
//
//// 要求出余额，我们需要得到未花费交易的输出
//// 交易输出这个结构体中包含 value 字段，求出所有未花费交易输出的 value 之和就是账户余额
//// 首先求得未花费交易输出
//// 通过找到未花费输出交易的集合，我们返回集合中的所有未花费的交易输出
//func (bc *BlockChain) FindUTXO(address string) []transaction.TXOutput {
//	var UTXOs []transaction.TXOutput
//	// 找到 address 地址下的未花费交易输出的交易的集合
//	unspentTransaction := bc.FindUnspentTransaction(address)
//	// 遍历交易集合的到交易，从交易中提取出输出字段 Vout，从输出字段中提取出属于 address 的输出
//	for _, tx := range unspentTransaction {
//		for _, out := range tx.Vout {
//			if out.CanBeUnlockedWith(address) {
//				UTXOs = append(UTXOs, out)
//			}
//		}
//	}
//
//	// 返回未花费交易输出
//	return UTXOs
//}
//
//// 发送币操作，相当于创建一笔未花费输出交易
//func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *transaction.Transaction {
//	var inputs []transaction.TXInput
//	var outputs []transaction.TXOutput
//
//	// validOutputs 是一个存放要用到的未花费输出的交易/输入的 map
//	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
//
//	if acc < amount {
//		log.Error("Error: Not enough tokens...")
//		return nil
//	}
//
//	// 通过 validOutputs 里面的数据来放入建立一个输入列表
//	for txid, outs := range validOutputs {
//		// 反序列化的到 txID
//		txID, err := hex.DecodeString(txid)
//		if err != nil {
//			log.Error("反序列化失败: ", err)
//			return nil
//		}
//		// 遍历输出 outs 切片，得到 TXInput 里的 Vout 字段
//		for _, out := range outs {
//			input := transaction.TXInput{
//				Txid:      txID,
//				Vout:      out,
//				ScriptSig: from,
//			}
//			inputs = append(inputs, input)
//		}
//	}
//	// 建立一个输出列表
//	outputs = append(outputs, transaction.TXOutput{amount, to})
//	if acc > amount {
//		// 相当于找零
//		outputs = append(outputs, transaction.TXOutput{
//			Value:        acc - amount,
//			ScriptPubKey: from,
//		})
//	}
//	tx := transaction.Transaction{
//		ID:   nil,
//		Vin:  inputs,
//		Vout: outputs,
//	}
//	tx.SetID()
//
//	return &tx
//}
//
//// 找到可以花费的交易输出，这是基于上面的 FindUnspentTransactions 方法
//func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
//	// 未花费交易输出 map 集合
//	unspentOutput := make(map[string][]int)
//	// 未花费交易
//	unspentTXs := bc.FindUnspentTransaction(address)
//	// 累加未花费交易输出中的 Value 值
//	accumulated := 0
//
//Work:
//	for _, tx := range unspentTXs {
//		txID := hex.EncodeToString(tx.ID)
//
//		for outIdx, out := range tx.Vout {
//			if out.CanBeUnlockedWith(address) && accumulated < amount {
//				accumulated += out.Value
//				unspentOutput[txID] = append(unspentOutput[txID], outIdx)
//
//				if accumulated >= amount {
//					break Work
//				}
//			}
//		}
//	}
//
//	return accumulated,unspentOutput
//}


/*
	区块链实现
*/
const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
//区块链
type Blockchain struct {
	tip		[]byte
	db 		*bolt.DB
}

//工厂模式db
func(bc *Blockchain) Db() *bolt.DB {
	return bc.db
}

//把区块添加进区块链,挖矿
func (bc *Blockchain) MineBlock(transactions []*transaction.Transaction) {
	var lastHash []byte
	//只读的方式浏览数据库，获取当前区块链顶端区块的哈希，为加入下一区块做准备
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))	//通过键"l"拿到区块链顶端区块哈希

		return nil
	})
	if err != nil {
		log.Error(err)
		return
	}

	//prevBlock := bc.Blocks[len(bc.Blocks)-1]
	//求出新区块
	newBlock := pow.NewBlock(transactions,lastHash)
	// bc.Blocks = append(bc.Blocks,newBlock)
	//把新区块加入到数据库区块链中
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		if err != nil {
			log.Error(err)
			return nil
		}
		err = b.Put([]byte("l"),newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

//创建创世区块  /修改/
func NewGenesisBlock(coinbase *transaction.Transaction) *block.Block {
	return pow.NewBlock([]*transaction.Transaction{coinbase},[]byte{})
}
//实例化一个区块链,默认存储了创世区块 ,接收一个地址为挖矿奖励地址 /修改/
func NewBlockchain(address string) *Blockchain {
	//return &Blockchain{[]*block.Block{NewGenesisBlock()}}
	var tip []byte
	//打开一个数据库文件，如果文件不存在则创建该名字的文件
	db,err := bolt.Open(dbFile,0600,nil)
	if err != nil {
		log.Error(err)
		return nil
	}
	//读写操作数据库
	err = db.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBucket))
		//查看名字为blocksBucket的Bucket是否存在
		if b == nil {
			//不存在则从头 创建
			fmt.Println("不存在区块链，需要重新创建一个区块链...")

			//genesis := NewGenesisBlock()	//创建创世区块
			//此时的创世区块就要包含交易coinbaseTx
			cbtx := transaction.NewCoinbaseTX(address,genesisCoinbaseData)
			genesis := NewGenesisBlock(cbtx)

			b,err := tx.CreateBucket([]byte(blocksBucket)) //创建名为blocksBucket的桶
			if err != nil {
				log.Error(err)
				return nil
			}
			err = b.Put(genesis.Hash,genesis.Serialize()) //写入键值对，区块哈希对应序列化后的区块
			if err != nil {
				log.Error(err)
				return nil
			}
			err = b.Put([]byte("l"),genesis.Hash) //"l"键对应区块链顶端区块的哈希
			if err != nil {
				log.Error(err)
				return nil
			}
			tip = genesis.Hash //指向最后一个区块，这里也就是创世区块
		} else {
			//如果存在blocksBucket桶，也就是存在区块链
			//通过键"l"映射出顶端区块的Hash值
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip,db}  //此时Blockchain结构体字段已经变成这样了
	return &bc
}

//分割线——————迭代器——————
type BlockchainIterator struct {
	currentHash 	[]byte
	db 				*bolt.DB
}
//当需要遍历当前区块链时，创建一个此区块链的迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip,bc.db}

	return bci
}

//迭代器的任务就是返回链中的下一个区块
func (i *BlockchainIterator) Next() *block.Block {
	var Block *block.Block

	//只读方式打开区块链数据库
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//获取数据库中当前区块哈希对应的被序列化后的区块
		encodeBlock := b.Get(i.currentHash)
		//反序列化，获得区块
		Block = block.DeserializeBlock(encodeBlock)

		return nil
	})
	if err != nil {
		log.Error(err)
		return nil
	}

	//把迭代器中的当前区块哈希设置为上一区块的哈希，实现迭代的作用
	i.currentHash =Block.PrevBlockHash

	return Block

}

//在区块链上找到每一个区块中属于address用户的未花费交易输出,返回未花费输出的交易切片
func (bc *Blockchain) FindUnspentTransactions(address string) []transaction.Transaction {
	var unspentTXs []transaction.Transaction
	//创建一个map，存储已经花费了的交易输出
	spentTXOs := make(map[string][]int)
	//因为要在链上遍历区块，所以要使用到迭代器
	bci := bc.Iterator()

	for {
		block := bci.Next()  //迭代

		//遍历当前区块上的交易
		for _,tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID) //把交易ID转换成string类型，方便存入map中

			//标签
		Outputs:
			//遍历当前交易中的输出切片，取出交易输出
			for outIdx,out := range tx.Vout {
				//在已经花费了的交易输出map中，如果没有找到对应的交易输出，则表示当前交易的输出未花费
				//反之如下
				if spentTXOs[txID] != nil {
					//存在当前交易的输出中有已经花费的交易输出，
					//则我们遍历map中保存的该交易ID对应的输出的index
					//提示：(这里的已经花费的交易输出index其实就是输入TXInput结构体中的Vout字段)
					for _,spentOut := range spentTXOs[txID] {
						//首先要清楚当前交易输出是一个切片，里面有很多输出，
						//如果map里存储的引用的输出和我们当前遍历到的输出index重合,则表示该输出被引用了
						if spentOut == outIdx {
							continue Outputs  //我们就继续遍历下一轮，找到未被引用的输出
						}
					}
				}
				//到这里是得到此交易输出切片中未被引用的输出

				//这里就要从这些未被引用的输出中筛选出属于该用户address地址的输出
				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs,*tx)
				}
			}
			//判断是否为coinbase交易
			if tx.IsCoinbase() == false {
				//如果不是,则遍历当前交易的输入
				for _,in := range tx.Vin {
					//如果当前交易的输入是被该地址address所花费的，就会有对应的该地址的引用输出
					//则在map上记录该输入引用的该地址对应的交易输出
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID],in.Vout)
					}
				}
			}
		}
		//退出for循环的条件就是遍历到的创世区块后
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return unspentTXs
}

//通过找到未花费输出交易的集合，我们返回集合中的所有未花费的交易输出
func (bc *Blockchain) FindUTXO(address string) []transaction.TXOutput {
	var UTXOs []transaction.TXOutput
	//找到address地址下的未花费交易输出的交易的集合
	unspentTransactions := bc.FindUnspentTransactions(address)
	//遍历交易集合得到交易，从交易中提取出输出字段Vout,从输出字段中提取出属于address的输出
	for _,tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs,out)
			}
		}
	}
	//返回未花费交易输出
	return UTXOs
}

//发送币操作,相当于创建一笔未花费输出交易
func NewUTXOTransaction(from,to string,amount int,bc *Blockchain) *transaction.Transaction {
	var inputs []transaction.TXInput
	var outputs []transaction.TXOutput
	//validOutputs是一个存放要用到的未花费输出的交易/输出的map
	acc,validOutputs := bc.FindSpendableOutputs(from,amount)

	if acc < amount {
		log.Error("ERROR:Not enough tokens...")
		return nil
	}
	//通过validOutputs里面的数据来放入建立一个输入列表
	for txid,outs := range validOutputs {
		//反序列化得到txID
		txID,err := hex.DecodeString(txid)
		if err != nil {
			log.Error(err)
			return nil
		}
		//遍历输出outs切片,得到TXInput里的Vout字段值
		for _,out := range outs {
			input := transaction.TXInput{txID,out,from}
			inputs = append(inputs,input)
		}
	}
	//建立一个输出列表
	outputs = append(outputs,transaction.TXOutput{amount,to})
	if acc > amount {
		outputs = append(outputs,transaction.TXOutput{acc - amount,from}) //相当于找零
	}
	tx := transaction.Transaction{nil,inputs,outputs}
	tx.SetID()

	return &tx
}

//找到可以花费的交易输出,这是基于上面的FindUnspentTransactions 方法
func (bc *Blockchain) FindSpendableOutputs(address string,amount int) (int,map[string][]int) {
	//未花费交易输出map集合
	unspentOutputs := make(map[string][]int)
	//未花费交易
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0	//累加未花费交易输出中的Value值

Work:
	for _,tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx,out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID],outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}
	return accumulated,unspentOutputs
}