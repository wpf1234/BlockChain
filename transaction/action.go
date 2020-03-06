package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	log "github.com/sirupsen/logrus"
)

//const subsidy = 50 //挖矿奖励
///*
//    创建一个交易的数据结构，交易是由交易 ID、交易输入、交易输出组成
//	一个交易有多个输入和多个输出，所以这里的交易输入和输出应该是切片类型
//*/
//type Transaction struct {
//	ID   []byte
//	Vin  []TXInput
//	Vout []TXOutput
//}
//
///*
//    1.每一笔交易的输入都会引用之前交易的一笔或多笔交易输出
//	2.交易输出保存了输入的值和锁定该输出的信息
//	3.交易输入保存了引用之前交易输出的交易ID、具体到引用该交易的第几个输出、
//	  能正确解锁引用输出的签名信息
//*/
//// 交易输出
//type TXOutput struct {
//	Value        int    // 输出的值(可以理解为金额)
//	ScriptPubKey string // 锁定该输出的脚本(目前还没有实现地址，所以还不能锁定该输出为具体哪个地址所有)
//}
//
//// 交易输入
//type TXInput struct {
//	Txid      []byte // 引用的事前交易的ID
//	Vout      int    // 引用之前交易输出的具体是哪一个输出(一个交易中输出一般有多个)
//	ScriptSig string // 能解锁引用输出交易的签名脚本(目前还没实现地址)
//}
//
///*
//	区块链上存储的交易都是由这些输入输出交易所组成的，
//一个输入交易必须引用之前的输出交易，一个输出交易会被之后的输入所引用。
//    问题来了，在最开始的区块链上是先有输入还是先有输出喃？
//答案是先有输出，因为是区块链的创世区块产生了第一个输出，
//这个输出也就是我们常说的挖矿奖励-狗头金，每一个区块都会有一个这样的输出，
//这是奖励给矿工的交易输出，这个输出是凭空产生的。
//*/
//// 现在我们来创建一个这样的 coinbase 挖矿输出
//// to 代表此输出奖励给谁，一般都是矿工地址，data 是交易附带的信息
//func NewCoinbaseTX(to, data string) *Transaction {
//	if data == "" {
//		data = fmt.Sprintf("奖励给 '%s' ", to)
//	}
//	// 此交易中的交易输入，没有交易输入信息
//	txin := TXInput{
//		Txid:      []byte{},
//		Vout:      -1,
//		ScriptSig: data,
//	}
//	// 交易输出，subsidy 为奖励矿工的币的数量
//	txout := TXOutput{
//		Value:        subsidy,
//		ScriptPubKey: to,
//	}
//	// 组成交易
//	tx := Transaction{
//		ID:   nil,
//		Vin:  []TXInput{txin},
//		Vout: []TXOutput{txout},
//	}
//	// 设置该交易的ID
//	tx.SetID()
//	return &tx
//}
//
//// 设置交易的ID，交易 ID 是序列化 TX后再哈希
//func (tx *Transaction) SetID() {
//	var hash [32]byte
//	var encoder bytes.Buffer
//
//	enc := gob.NewEncoder(&encoder)
//	err := enc.Encode(tx)
//	if err != nil {
//		log.Error("序列化失败: ", err)
//		return
//	}
//	hash = sha256.Sum256(encoder.Bytes())
//	tx.ID = hash[:]
//}
//
///*
//	1.每一个区块至少存储一笔交易 coinbase 交易，所以我们在区块的字段中把 Data 字段换成交易
//	2.把所有涉及之前 Data 字段都要换，如 NewBlock(),GenesisBlock(),pow里的函数
//*/
//
//// 定义在输入和输出上的锁定和解锁方法，目的是让用户只能花自己所用于地址上的币
//// 输入上锁定的秘钥，表示能引用的输出是 unlockingData
//func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
//	return in.ScriptSig == unlockingData
//}
//
//// 输出上的解锁秘钥，表示能被引用的输入是 unlockingData
//func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
//	return out.ScriptPubKey == unlockingData
//}
//
//// 判断是否为 coinbase 交易
//func (tx *Transaction) IsCoinbase() bool{
//	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
//}

var subsidy int = 50  //挖矿奖励
/*创建一个交易的数据结构，交易是由交易ID、交易输入、交易输出组成的,
一个交易有多个输入和多个输出，所以这里的交易输入和输出应该是切片类型的
*/
type Transaction struct {
	ID		[]byte
	Vin		[]TXInput
	Vout	[]TXOutput
}

/*
1、每一笔交易的输入都会引用之前交易的一笔或多笔交易输出
2、交易输出保存了输出的值和锁定该输出的信息
3、交易输入保存了引用之前交易输出的交易ID、具体到引用
该交易的第几个输出、能正确解锁引用输出的签名信息
*/
//交易输出
type TXOutput struct {
	Value			int	//输出的值（可以理解为金额）
	ScriptPubKey	string	// 锁定该输出的脚本（目前还没实现地址，所以还不能锁定该输出为具体哪个地址所有）
}
//交易输入
type TXInput struct {
	Txid 	[]byte //引用的之前交易的ID
	Vout	int 	//引用之前交易输出的具体是哪个输出（一个交易中输出一般有很多）
	ScriptSig	string  // 能解锁引用输出交易的签名脚本（目前还没实现地址，所以本章不能实现此功能）
}

/*
	区块链上存储的交易都是由这些输入输出交易所组成的，
一个输入交易必须引用之前的输出交易，一个输出交易会被之后的输入所引用。
    问题来了，在最开始的区块链上是先有输入还是先有输出喃？
答案是先有输出，因为是区块链的创世区块产生了第一个输出，
这个输出也就是我们常说的挖矿奖励-狗头金，每一个区块都会有一个这样的输出，
这是奖励给矿工的交易输出，这个输出是凭空产生的。
*/
//现在我们来创建一个这样的coinbase挖矿输出
//to 代表此输出奖励给谁，一般都是矿工地址，data是交易附带的信息
func NewCoinbaseTX(to,data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("奖励给 '%s'",to)
	}
	//此交易中的交易输入,没有交易输入信息
	txin := TXInput{[]byte{},-1,data}
	//交易输出,subsidy为奖励矿工的币的数量
	txout := TXOutput{subsidy,to}
	//组成交易
	tx := Transaction{nil,[]TXInput{txin},[]TXOutput{txout}}
	//设置该交易的ID
	tx.SetID()
	return &tx
}
//设置交易ID，交易ID是序列化tx后再哈希
func (tx *Transaction) SetID() {
	var hash [32]byte
	var encoder bytes.Buffer

	enc := gob.NewEncoder(&encoder)
	err := enc.Encode(tx)
	if err != nil {
		log.Error(err)
		return
	}
	hash = sha256.Sum256(encoder.Bytes())
	tx.ID =  hash[:]
}

/*
1、每一个区块至少存储一笔coinbase交易，所以我们在区块的字段中把Data字段换成交易。
2、把所有涉及之前Data字段都要换了，比如NewBlock()、GenesisBlock()、pow里的函数
*/

//定义在输入和输出上的锁定和解锁方法，目的是让用户只能花自己所用于地址上的币
//输入上锁定的秘钥,表示能引用的输出是unlockingData
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
//输出上的解锁秘钥,表示能被引用的输入是unlockingData
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

//判断是否为coinbase交易
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}
