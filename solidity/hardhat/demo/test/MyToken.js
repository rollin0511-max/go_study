const {expect} = require("chai");
const hre = require("hardhat");

describe("MyToken Test", async() => {

    const {ethers} = hre  // js 获取 hardhat对象
    const initialSupply = 10000  // 初始默认铸币数
    let MyTokenContract  // 合约实例
    let account1,account2  // 声明账号
    // async 异步函数
    // await 等待异步函数执行完成
    beforeEach(async () => {
        // console.log("等待2s")
        // await new Promise((resolve) => {
        //     setTimeout(() => {
        //         resolve(1)
        //     }, 2000);
        // })
        // console.log("开始运行测试用例")

        [account1,account2] = await ethers.getSigners()
        console.log("accounts:",account1,account2)

        const MyToken = await ethers.getContractFactory("MyToken")
        MyTokenContract = await MyToken.connect(account2).deploy(initialSupply)
        MyTokenContract.waitForDeployment()
        const contractAddress = await MyTokenContract.getAddress()
        expect(contractAddress).to.length.greaterThan(0)
        console.log("合约部署地址:",contractAddress)
    })

    it("验证下合约的 name symbol decimal",async  () => {
        const name = await MyTokenContract.name()
        expect(name).to.equal("MyToken")

        const symbol = await MyTokenContract.symbol()
        expect(symbol).to.equal("MTK")

        const decimal = await MyTokenContract.decimals()
        console.log("合约的 decimal:",decimal)
        expect(decimal).to.equal(18)
    })

    it("测试转账",async  () => {
        // const balanceOfAccount1 = await MyTokenContract.balanceOf(account1.address)
        // expect(balanceOfAccount1).to.equal(initialSupply)
        // console.log("account1 余额:",balanceOfAccount1)
        const resp = await MyTokenContract.transfer(account1,initialSupply/2)
        console.log("转账交易:",resp)
        const balanceOfAccount2 = await MyTokenContract.balanceOf(account2.address)
        expect(balanceOfAccount2).to.equal(initialSupply/2)
        console.log("account2 余额:",balanceOfAccount2)
    })
})

