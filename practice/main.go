package main

import "fmt"

// PayMethod 定义一个支付接口
type PayMethod interface {
	// Pay 公有的支付方法
	Pay(int)
}

// CreditCard 定义一个信用卡结构体
type CreditCard struct {
	// balance 余额
	balance int
	// limit 限额
	limit int
}

// Pay 实现支付方法
func (c *CreditCard) Pay(amout int) {
	if c.balance < amout {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amout
}

// anyParam 任意参数
func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

// main 主函数
func main() {
	// 定义一个余额为100 限额为1000的信用卡
	c := CreditCard{balance: 1500, limit: 1000}
	// 调用支付方法
	c.Pay(200)
	// 打印余额
	fmt.Println("c.balance: ", c.balance)
	// 打印限额
	fmt.Println("c.limit: ", c.limit)
	// 定义一个支付接口
	var a PayMethod = &c
	fmt.Println("a.Pay: ", a)

	var b interface{} = &c
	fmt.Println("b: ", b)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
}
