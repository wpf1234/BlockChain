package CLI

import (
	"BlockChain/blockchain"
	"BlockChain/pow"
	"BlockChain/transaction"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

//// 首先我们想要拥有这些命令
//// 1.加入区块命令
//// 2.打印区块链命令
//
//// 创建一个 CLI 结构体
//
//type CLI struct {
//	//BC *blockchain.BlockChain
//}
//
////入口函数
//
////func (cli *CLI) Run() {
////	//判断命令行输入参数的个数，如果没有输入任何参数则打印提示输入参数信息
////	cli.validateArgs()
////	//实例化flag集合
////	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
////	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
////
////	//注册一个flag标志符
////	addBlockData := addBlockCmd.String("data", " ", "区块数据")
////
////	switch os.Args[1] { //os.Args为一个保存输入命令的切片
////	case "addblock":
////		//解析出"addblock"后面的命令
////		err := addBlockCmd.Parse(os.Args[2:])
////		if err != nil {
////			log.Error(err)
////			return
////		}
////	case "printchain":
////		err := printChainCmd.Parse(os.Args[2:])
////		if err != nil {
////			log.Error(err)
////			return
////		}
////	default:
////		cli.printUsage() //提示用户怎么正确输入命令
////		os.Exit(1)
////	}
////
////	//进入被解析出的命令，进一步操作
////	if addBlockCmd.Parsed() {
////		if *addBlockData == " " {
////			addBlockCmd.Usage() //如果没有输入标志位data，则提醒用户怎么正确的输入
////			os.Exit(1)
////		}
////		//用户输入正确则进行加入区块的操作
////		cli.addBlock(*addBlockData)
////	}
////	if printChainCmd.Parsed() {
////		//打印区块链操作
////		cli.printChain()
////	}
////}
//func (cli *CLI) Run() {
//	//判断命令行输入参数的个数，如果没有输入任何参数则打印提示输入参数信息
//	cli.validateArgs()
//	//实例化flag集合
//	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
//	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
//	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
//	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
//	//注册flag标志符
//	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
//	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
//	sendFrom := sendCmd.String("from", "", "Source wallet address")
//	sendTo := sendCmd.String("to", "", "Destination wallet address")
//	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
//
//	switch os.Args[1] {		//os.Args为一个保存输入命令的切片
//	case "getbalance":
//		err := getBalanceCmd.Parse(os.Args[2:])
//		if err != nil {
//			log.Error(err)
//		}
//	case "createblockchain":
//		err := createBlockchainCmd.Parse(os.Args[2:])
//		if err != nil {
//			log.Error(err)
//		}
//	case "printchain":
//		err := printChainCmd.Parse(os.Args[2:])
//		if err != nil {
//			log.Error(err)
//		}
//	case "send":
//		err := sendCmd.Parse(os.Args[2:])
//		if err != nil {
//			log.Error(err)
//		}
//	default:
//		cli.printUsage()
//		os.Exit(1)
//	}
//
//	//进入被解析出的命令，进一步操作
//	if getBalanceCmd.Parsed() {
//		if *getBalanceAddress == "" {
//			getBalanceCmd.Usage()
//			os.Exit(1)
//		}
//		cli.getBalance(*getBalanceAddress)
//	}
//
//	if createBlockchainCmd.Parsed() {
//		if *createBlockchainAddress == "" {
//			createBlockchainCmd.Usage()
//			os.Exit(1)
//		}
//		cli.createBlockchain(*createBlockchainAddress)
//	}
//
//	if printChainCmd.Parsed() {
//		cli.printChain()
//	}
//
//	if sendCmd.Parsed() {
//		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
//			sendCmd.Usage()
//			os.Exit(1)
//		}
//		cli.send(*sendFrom, *sendTo, *sendAmount)
//	}
//}
//
////加入输入格式错误信息提示
//func (cli *CLI) printUsage() {
//	fmt.Println("Usage:")
//	//fmt.Println("  addblock -data  区块信息")
//	//fmt.Println("  printchain - Print all the blocks of the blockchain")
//	fmt.Println("  getbalance -address ADDRESS  得到该地址的余额")
//	fmt.Println("  createblockchain -address ADDRESS 创建一条链并且该地址会得到狗头金")
//	fmt.Println("  printchain - 打印链")
//	fmt.Println("  send -from FROM -to TO -amount AMOUNT 地址from发送amount的币给地址to")
//
//}
//
////判断命令行参数，如果没有输入参数则显示提示信息
//func (cli *CLI) validateArgs() {
//	if len(os.Args) < 2 {
//		cli.printUsage()
//		os.Exit(1)
//	}
//}
//
////加入区块函数调用
////func (cli *CLI) addBlock(data string) {
////	cli.BC.AddBlock(data)
////	fmt.Println("成功加入区块...")
////}
//
////打印区块链函数调用
//func (cli *CLI) printChain() {
//	//这里需要用到迭代区块链的思想
//	//创建一个迭代器
//	//bci := cli.BC.Iterator()
//	//
//	//for {
//	//
//	//	block := bci.Next() //从顶端区块向前面的区块迭代
//	//
//	//	fmt.Printf("PrevHash:%x\n", block.PrevBlockHash)
//	//	fmt.Printf("Data:%s\n", block.Data)
//	//	fmt.Printf("Hash:%x\n", block.Hash)
//	//	//验证当前区块的pow
//	//	pow := pow.NewProofOfWork(block)
//	//	boolen := pow.Validate()
//	//	fmt.Printf("POW is %s\n", strconv.FormatBool(boolen))
//	//	fmt.Println()
//	//
//	//	if len(block.PrevBlockHash) == 0 {
//	//		break
//	//	}
//	//}
//
//	// 实例化一条链
//	// 因为已经有链了，不会重新创建链，所以接收的 address 设置为空
//	bc:=blockchain.NewBlockChain("")
//	defer bc.DB().Close()
//
//	// 这里需要用到迭代区块链的思想
//	// 创将一个迭代器
//	bci:=bc.Iterator()
//
//	for{
//		// 从顶端区块向前面的区块迭代
//		block:=bci.Next()
//
//		fmt.Printf("时间戳:%v\n",block.TimeStamp)
//		fmt.Printf("前一个区块的哈希:%x\n",block.PrevBlockHash)
//		fmt.Printf("当前区块的哈希:%x\n",block.Hash)
//		// 验证当前区块的 pow
//		pow:=pow.NewProofOfWork(block)
//
//		boolen:=pow.Validate()
//		fmt.Printf("POW is %s\n",strconv.FormatBool(boolen))
//		fmt.Println()
//
//		if len(block.PrevBlockHash) == 0{
//			break
//		}
//	}
//}
//
//// 创建一条链
//func (cli *CLI) createBlockchain(address string){
//	bc:=blockchain.NewBlockChain(address)
//	bc.DB().Close()
//	fmt.Println("Done")
//}
//
//// 求账户余额
//func (cli *CLI) getBalance(address string){
//	bc:=blockchain.NewBlockChain(address)
//	defer bc.DB().Close()
//
//	balance:=0
//	UTXOs:=bc.FindUTXO(address)
//
//	// 遍历 UTXOs 中的交易输出 out，得到输出字段 out，value，求出余额
//	for _,out:=range UTXOs{
//		balance+=out.Value
//	}
//
//	fmt.Printf("Balance of '%s':%d\n",address,balance)
//}
//
//// send 方法
//func (cli *CLI) send(from,to string,amount int){
//	bc:=blockchain.NewBlockChain(from)
//
//	defer bc.DB().Close()
//
//	tx:=blockchain.NewUTXOTransaction(from,to,amount,bc)
//	// 挖出一个包含该交易的区块，此时区块只有一个交易
//	bc.MineBlock([]*transaction.Transaction{tx})
//	fmt.Println("发送成功......")
//}

//首先我们想要拥有这些命令 1.加入区块命令 2.打印区块链命令

//创建一个CLI结构体
type CLI struct {
	//BC *blockchain.Blockchain
}


//加入输入格式错误信息提示
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  getbalance -address ADDRESS  得到该地址的余额")
	fmt.Println("  createblockchain -address ADDRESS 创建一条链并且该地址会得到狗头金")
	fmt.Println("  printchain - 打印链")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT 地址from发送amount的币给地址to")
}

//判断命令行参数，如果没有输入参数则显示提示信息
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// //加入区块函数调用
// func (cli *CLI) addBlock(data string) {
// 	cli.BC.MineBlock(data)
// 	fmt.Println("成功加入区块...")
// }

//打印区块链函数调用
func (cli *CLI) printChain() {
	//实例化一条链
	bc := blockchain.NewBlockchain("")  //因为已经有了链，不会重新创建链，所以接收的address设置为空
	defer bc.Db().Close()

	//这里需要用到迭代区块链的思想
	//创建一个迭代器
	bci := bc.Iterator()

	for {

		block := bci.Next()	//从顶端区块向前面的区块迭代

		fmt.Printf("时间戳:%v\n",block.Timestamp)
		fmt.Printf("PrevHash:%x\n",block.PrevBlockHash)
		//fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		//验证当前区块的pow
		pow := pow.NewProofOfWork(block)
		boolen := pow.Validate()
		fmt.Printf("POW is %s\n",strconv.FormatBool(boolen))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

//创建一条链
func (cli *CLI) createBlockchain(address string) {
	bc := blockchain.NewBlockchain(address)
	bc.Db().Close()
	fmt.Println("Done!")
}
//求账户余额
func (cli *CLI) getBalance(address string) {
	bc := blockchain.NewBlockchain(address)
	defer bc.Db().Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	//遍历UTXOs中的交易输出out，得到输出字段out.Value,求出余额
	for _,out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s':%d\n",address,balance)
}

//send方法
func (cli *CLI) send(from,to string,amount int) {
	bc := blockchain.NewBlockchain(from)
	defer bc.Db().Close()

	tx := blockchain.NewUTXOTransaction(from,to,amount,bc)
	//挖出一个包含该交易的区块,此时区块只有这一个交易
	bc.MineBlock([]*transaction.Transaction{tx})
	fmt.Println("发送成功...")
}

//入口函数
func (cli *CLI) Run() {
	//判断命令行输入参数的个数，如果没有输入任何参数则打印提示输入参数信息
	cli.validateArgs()
	//实例化flag集合
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	//注册flag标志符
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {		//os.Args为一个保存输入命令的切片
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Error(err)
			return
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Error(err)
			return
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Error(err)
			return
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Error(err)
			return
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	//进入被解析出的命令，进一步操作
	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}
