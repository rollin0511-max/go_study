package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// 查询区块信息
	queryBlock()

	// 查询交易信息
	queryTransaction()

	// 查询收据
	queryReceipt()

	// 创建新钱包
	createWallet()

	// ETH转账
	transferETH()

	// 代币转账
	transferToken()

	// 查询余额
	queryBalance()

	// 查询代币余额
	queryTokenBalance()

	// 订阅区块
	subscribeBlock()
}

/**
 * 查询区块信息
 */
func queryBlock() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 5671744
	fmt.Println(header.Time)                // 1712798400
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))   // 70
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70
}

/**
 * 查询交易信息
 */
func queryTransaction() {
	// 获取客户端
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取区块信息
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 遍历交易
	for _, tx := range block.Transactions() {
		// 打印交易hash的Hex
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 打印交易金额的String
		fmt.Println(tx.Value().String()) // 100000000000000000
		// 打印交易gas的Uint64
		fmt.Println(tx.Gas()) // 21000
		// 打印交易gas价格的Uint64
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		// 打印交易nonce的Uint64
		fmt.Println(tx.Nonce()) // 245132
		// 打印交易数据的Hex
		fmt.Println(tx.Data()) // []
		// 打印交易接收地址的Hex
		fmt.Println(tx.To().Hex()) // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			// 打印交易发送地址的Hex
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		// 获取交易收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		// 打印交易收据的Status
		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	// 根据区块地址获取区块hash
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	// 获取区块交易数量
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	// 遍历区块交易
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}
	// 根据交易hash获取交易
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)
}

/**
 * 查询收据
 */
func queryReceipt() {
	// 获取客户端
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	// 指定区块获取区块hash
	blockNumber := big.NewInt(5671744)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	// 使用区块hash 调用 BlockReceipts 方法就可以得到指定区块中所有的收据列表
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	// 使用区块号 调用 BlockReceipts 方法就可以得到指定区块中所有的收据列表
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true

	// 遍历收据列表
	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}

	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)                // 1
	fmt.Println(receipt.Logs)                  // []
	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	fmt.Println(receipt.TransactionIndex)      // 0
	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
}

/**
 * 创建新钱包
 */
func createWallet() {
	// 导入 go-ethereum crypto 包，该包提供用于生成随机私钥的 GenerateKey 方法
	privateKey, err := crypto.GenerateKey()
	// 如果已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥
	// privateKey, err := crypto.HexToECDSA("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")
	if err != nil {
		log.Fatal(err)
	}
	// 导入 golang crypto/ecdsa 包并使用 FromECDSA 方法将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'  用于签署交易的私钥

	// 从私钥中提取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04' 用于展示的公钥

	// 从公钥中提取公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

/**
 * ETH转账
 */
func transferETH() {
	// 获取客户端连接
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	// 从私钥中提取公钥
	publicKey := privateKey.Public()
	// 将公钥转换为ECDSA公钥指针
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 从ECDSA公钥中提取公共地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 自定义转账金额
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	// 自定义转账Gas限制
	gasLimit := uint64(21000) // in units
	// 计算转账Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 从钱包地址字符串中获取目标地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	// 自定义转账交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	// 获取当前网络ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 接受一个未签名的事务和我们之前构造的私钥
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 将已签名的事务广播到整个网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

/**
 * 代币转账
 */
func transferToken() {
	// 获取客户端连接
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/")
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 从私钥中提取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 从公钥ECDSA指针中提取公共地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 代币传输不需要传输 ETH，因此将交易“值”设置为“0”
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 发送代币的地址存储在变量
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 将代币合约地址分配给变量
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	// 函数名将是传递函数的名称，即 ERC-20 规范中的 transfer 和参数类型。
	//	第一个参数类型是 address（令牌的接收者），第二个类型是 uint256（要发送的代币数量）。
	//	不需要没有空格和参数名称。 我们还需要用字节切片格式
	transferFnSignature := []byte("transfer(address,uint256)")
	// 生成函数签名的 Keccak256 哈希
	hash := sha3.NewLegacyKeccak256()
	// 写入函数签名到哈希
	hash.Write(transferFnSignature)
	// 从哈希中获取前4个字节，即方法ID
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	// 将给我们发送代币的地址左填充到 32 字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// 确定要发送多少个代币
	//	在这种情况下，我们将发送 1000 个代币。 我们需要将这个数字转换为一个 32 字节的大整数，
	//	并将其左填充到 32 字节。 我们将使用 big.Int 类型来完成此操作。
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// 将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 指定数据和地址 返回我们估算的完成交易所需的估计燃气上限
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	// 创建一个新的交易
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	// 获取当前网络ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 使用发件人的私钥对事务进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}

/**
 * 查询余额
 */
func queryBalance() {
	// 获取客户端连接
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	// // 调用 ethclient 的 BalanceAt 方法，给它传递账户地址和可选的区块号。将区块号设置为 nil 将返回最新的余额
	// 获取账户地址
	account := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	// 查询当前余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 传区块高度能读取指定区块时的账户余额
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	// 将余额转换为 ETH 值
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	// 查询当前待处理的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}

/**
 * 查询代币余额
 */
func queryTokenBalance() {
	//// 获取客户端连接
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 获取代币合约地址
	//// Golem (GNT) Address
	//tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	//// 创建代币合约实例
	//instance, err := token.NewToken(tokenAddress, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 获取钱包地址
	//address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	//// 查询代币余额
	//bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//name, err := instance.Name(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//symbol, err := instance.Symbol(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//decimals, err := instance.Decimals(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	//fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	//fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	//fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	//fbal := new(big.Float)
	//fbal.SetString(bal.String())
	//value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	//fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

/**
 * 订阅区块
 */
func subscribeBlock() {
	// 获取websocket RPC URL 客户端链接
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个新的通道，用于接收区块头
	headers := make(chan *types.Header)
	// 接收我们刚创建的区块头通道，该方法将返回一个订阅对象
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	// 订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			// 获取区块完整信息
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
