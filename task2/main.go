package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1.编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func addTen(num *int) {
	*num += 10
}

// 2.实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func double(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

// 3.编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func gorutineUse() {
	// 等待协程执行完成
	wg := sync.WaitGroup{}
	// 设置等待2个协程
	wg.Add(2)

	// 协程1: 打印奇数
	go func() {
		// 协程完成时通知
		defer wg.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("gorutine 奇数:", i)
			}
		}
	}()

	// 协程2: 打印偶数
	go func() {
		// 协程完成时通知
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("gorutine 偶数:", i)
			}
		}
	}()

	// 等待所有协程执行完成
	wg.Wait()
	fmt.Println("所有协程执行完成")
}

// 4.设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func schedule(tasks ...func()) []int {
	// 等待所有任务执行完成
	wg := sync.WaitGroup{}
	// 设置等待任务数量
	wg.Add(len(tasks))
	// 定义任务执行时间统计
	taskTimeSeconds := make([]int, len(tasks))
	// 并发执行任务
	for i, task := range tasks {
		go func(t func(), i int) {
			// 任务完成时通知
			defer wg.Done()
			// 执行任务
			start := time.Now()
			t()
			end := time.Now()
			// 计算任务执行时间
			taskTimeSeconds[i] = int(end.Sub(start).Seconds())
		}(task, i)
	}
	// 等待所有任务执行完成
	wg.Wait()
	fmt.Println("所有任务执行完成")
	return taskTimeSeconds
}

// Shape 5.定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area() string
	Perimeter() string
}

type Rectangle struct {
	height int
	width  int
}
type Circle struct {
	radius int
}

func (r Rectangle) Area() string {
	return fmt.Sprintf("矩形面积: %d", r.height*r.width)
}
func (r Rectangle) Perimeter() string {
	return fmt.Sprintf("矩形周长: %d", 2*(r.height+r.width))
}
func (c Circle) Area() string {
	return fmt.Sprintf("圆面积: %d", c.radius*c.radius)
}
func (c Circle) Perimeter() string {
	return fmt.Sprintf("圆周长: %d", 2*c.radius)
}

// 6.使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工ID: %d\n", e.Name, e.Age, e.EmployeeID)
}

// 7.编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func waitGroupUse() {
	// 等待所有任务执行完成
	wg := sync.WaitGroup{}
	// 设置等待任务数量
	wg.Add(2)
	// 定义一个通道
	channel := make(chan int)
	// 协程1: 生成整数
	go func() {
		// 协程完成时通知
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			channel <- i
			fmt.Println("发送到通道中的整数:", i)
			time.Sleep(1 * time.Second)
		}
		// 关闭通道
		close(channel)
	}()
	// 协程2: 接收通道中的整数并打印
	go func(ch <-chan int) {
		// 协程完成时通知
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收到并消费通道中的整数:", num)
		}
	}(channel)
	// 等待所有协程执行完成
	wg.Wait()
	fmt.Println("所有协程执行完成")
}

// 8.缓冲通道
func bufferChannel() {
	// 定义一个缓冲通道
	channel := make(chan int, 10)
	// 等待所有协程执行完成
	wg := sync.WaitGroup{}
	// 设置等待2个协程
	wg.Add(2)
	// 协程1: 发送整数到通道
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			channel <- i
			fmt.Println("发送到通道中的整数:", i)
		}
		close(channel)
	}()
	// 协程2: 接收通道中的整数并打印
	go func(channel <-chan int) {
		// 协程完成时通知
		defer wg.Done()
		for num := range channel {
			fmt.Println("接收到并消费通道中的整数:", num)
		}
	}(channel)
	// 等待所有协程执行完成
	wg.Wait()
	fmt.Println("所有协程执行完成")
}

// 9.编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func counterUse() {
	// 定义一个共享的计数器
	counter := 0
	// 定义一个互斥锁
	mutex := sync.Mutex{}
	// 等待所有协程执行完成
	wg := sync.WaitGroup{}
	// 设置等待10个协程
	wg.Add(10)
	// 启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			// 协程完成时通知
			defer wg.Done()
			// 对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 加锁
				mutex.Lock()
				// 递增计数器
				counter++
				fmt.Println("计数器当前值counter:", counter)
				// 解锁
				mutex.Unlock()
			}
		}()
	}
	// 等待所有协程执行完成
	wg.Wait()
	// 输出计数器的值
	fmt.Println("计数器的值:", counter)
}

// 10. 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func counterUseAtomic() {
	// 定义一个共享的计数器
	var counter int64
	// 等待所有协程执行完成
	wg := sync.WaitGroup{}
	// 设置等待10个协程
	wg.Add(10)
	// 启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			// 协程完成时通知
			defer wg.Done()
			// 对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 原子递增计数器
				atomic.AddInt64(&counter, 1)
				fmt.Println("原子操作计数器当前值:", counter)
			}
		}()
	}
	// 等待所有协程执行完成
	wg.Wait()
	// 输出计数器的值
	fmt.Println("原子操作计数器的值:", counter)
}

func main() {
	// 1. addTen
	// 定义一个整数
	//num := 1
	//// 调用函数
	//addTen(&num)
	//// 输出修改后的值
	//fmt.Println("num: ", num)

	// 2. double
	// 定义一个整数切片
	//slice := []int{1, 2, 3, 4, 5}
	//fmt.Println("double前 slice: ", slice)
	//// 调用函数
	//double(&slice)
	//// 输出修改后的值
	//fmt.Println("double后 slice: ", slice)

	// 3. gorutineUse
	// gorutineUse()

	// 4. schedule
	//tasks := []func(){
	//	func() {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("任务1")
	//	},
	//	func() {
	//		time.Sleep(3 * time.Second)
	//		fmt.Println("任务2")
	//	},
	//}
	//taskTimeSeconds := schedule(tasks...)
	//// 输出任务执行时间
	//for i, t := range taskTimeSeconds {
	//	fmt.Printf("任务%d执行时间: %d秒\n", i+1, t)
	//}

	// 5.调用Shape接口
	//rect := Rectangle{
	//	height: 10,
	//	width:  20,
	//}
	//circle := Circle{
	//	radius: 5,
	//}
	//fmt.Println(rect.Area())
	//fmt.Println(rect.Perimeter())
	//fmt.Println(circle.Area())
	//fmt.Println(circle.Perimeter())

	//6. 打印员工信息
	//emp := Employee{
	//	Person: Person{
	//		Name: "张三",
	//		Age:  30,
	//	},
	//	EmployeeID: 1001,
	//}
	//emp.PrintInfo()

	// 7. 通道使用
	//waitGroupUse()
	// 8.缓冲通道
	//bufferChannel()
	// 9. 计数器
	//counterUse()
	// 10. 无锁原子计数器
	counterUseAtomic()
}
