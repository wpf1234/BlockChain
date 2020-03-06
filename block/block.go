package block

import (
	"BlockChain/transaction"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	log "github.com/sirupsen/logrus"
)

//// 创建一个区块
//type Block struct {
//	TimeStamp int64 // 时间戳
//	//Data          []byte // 区块中所承载的数据
//	Transactions  []*transaction.Transaction // 将原来的 Data 字段换为交易数组
//	PrevBlockHash []byte                     // 上一区块的哈希值
//	Hash          []byte                     // 本区块的哈希值
//	Nonce         int                        // 一个随机数值
//}
//
///*
//   在 pow 算法中，以下函数不使用
//*/
//
//// 声明一个求本区块的 Hash 值的方法
//// 求区块哈希值是把本区块的时间戳、区块数据、上一区块 Hash 值一起做 sha256 哈希处理
////func (b *Block) SetHash() {
////	// 把时间戳转换成十进制再强制转换成 []byte 类型
////	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
////	// 将一系列 []byte 切片链接为一个 []byte 切片
////	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
////	hash := sha256.Sum256(header)
////
////	b.Hash = hash[:]
////}
//
//// 创建一个区块
////func NewBlock(data string, prevBlockHash []byte) *Block {
////	block := &Block{
////		TimeStamp:     time.Now().Unix(),
////		Data:          []byte(data),
////		PrevBlockHash: prevBlockHash,
////		Hash:          []byte{},
////	}
////	block.SetHash() // 给 block.Hash 赋值
////	return block
////}
//
//// 实现 Block 的序列化
//// gob 是 golang 自带的一个主句结构序列化的编码/解码工具
//// 由发送端使用 Encoder 对数据进行编码，接收端收到信息后，用 Decoder 进行解码
//func (b *Block) Serialize() []byte {
//	// 首先定义一个 Buffer 存储序列化后的数据
//	var result bytes.Buffer
//	// 实例化一个序列化实例，结果保存到 result 中
//	encoder := gob.NewEncoder(&result)
//	// 对区块进行实例化
//	err := encoder.Encode(b)
//	if err != nil {
//		log.Error(err)
//		return nil
//	}
//	return result.Bytes()
//}
//
//// 实现反序列化
//func DeserializeBlock(d []byte) *Block {
//	var block Block
//	decoder := gob.NewDecoder(bytes.NewReader(d))
//	err := decoder.Decode(&block)
//	if err != nil {
//		log.Error(err)
//		return nil
//	}
//
//	return &block
//}
//
//// 区块交易字段的哈希，因为每个交易 ID 是序列化并哈希后的交易数据，所以我们只需要把所有交易的 ID 进行哈希即可
//func (b * Block) HashTransactions() []byte{
//	var txHash [32]byte
//	var txHashes [][]byte
//	for _,tx:=range b.Transactions{
//		txHashes=append(txHashes,tx.ID)
//	}
//
//	txHash=sha256.Sum256(bytes.Join(txHashes,[]byte{}))
//
//	return txHash[:]
//}


//区块的结构体
type Block struct {
	Timestamp		int64
	Transactions	[]*transaction.Transaction
	PrevBlockHash	[]byte
	Hash 			[]byte
	Nonce			int
}

//区块交易字段的哈希,因为每个交易ID是序列化并哈希后的交易数据结构，所以我们只需把所以交易的ID进行哈希就可以了
func (b *Block) HashTransactions() []byte {
	var txHash [32]byte
	var txHashes [][]byte
	for _,tx := range b.Transactions {
		txHashes = append(txHashes,tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))

	return txHash[:]
}

//0.3 实现Block的序列化
func (b *Block) Serialize() []byte {
	//首先定义一个buffer存储序列化后的数据
	var result bytes.Buffer
	//实例化一个序列化实例,结果保存到result中
	encoder := gob.NewEncoder(&result)
	//对区块进行实例化
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//0.3 实现反序列化函数
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
