package main

import "BlockChain/CLI"

func main(){
	// 创建一条区块链
	//bc:=blockchain.NewBlockChain()

	/*
	    简单实现区块和区块链
	 */
	//bc.AddBlock("Send 50.0 BTC to Minner01")
	//time.Sleep(2*time.Second) //延时记入下一区块，让时间戳不同
	//bc.AddBlock("Send 25.0 BTC to Minner02")
	//
	////遍历区块链
	//for _,block := range bc.Blocks() {
	//	fmt.Printf("Prev.hash:%x\n",block.PrevBlockHash)
	//	fmt.Printf("Data:%s\n",block.Data)
	//	fmt.Printf("TimeStamp:%d\n",block.TimeStamp)
	//	fmt.Printf("Hash:%x\n",block.Hash)
	//	fmt.Println()
	//}

	/*
	    POW 算法下实现区块和区块链
	 */
	//bc.AddBlock("区块01")
	//bc.AddBlock("区块02")
	//
	//// 打印出区块链中各个区块的信息，并验证各个区块是否合格
	//for _,b:=range bc.Blocks(){
	//	fmt.Printf("时间戳：%v\n",b.TimeStamp)
	//	fmt.Printf("承载数据：%s\n",b.Data)
	//	fmt.Printf("上一区块哈希：%x\n",b.PrevBlockHash)
	//	fmt.Printf("本区块哈希：%x\n",b.Hash)
	//	fmt.Printf("随机数值：%v\n",b.Nonce)
	//	//验证当前区块的pow
	//	pow := pow.NewProofOfWork(b)
	//	boolen := pow.Validate()
	//	fmt.Printf("POW is %s\n",strconv.FormatBool(boolen))
	//	fmt.Println()
	//
	//}

	// 加入数据库，使用BlotDB
	// BlotDB 是使用键值对映射的方式来存储数据的，这些键值对被存储在一个叫bucket（桶）中

	// 这里 bc 中的字段 db 由于是小写字母开头，所以工厂模式由函数 DB() 返回 db
	// 程序退出前关闭数据库
	//defer bc.DB().Close()
	//
	//cli:=CLI.CLI{
	//	BC:bc,
	//}
	//cli.Run()

	// 添加交易
	cli:=CLI.CLI{}
	cli.Run()
}