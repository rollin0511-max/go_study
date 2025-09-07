package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义学生结构体
type student struct {
	gorm.Model
	Name  string `gorm:"size:255;comment:学生姓名"`
	Age   int    `gorm:"comment:学生年龄"`
	Grade string `gorm:"size:255;comment:学生年级"`
}

// 题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

// 定义account表
type account struct {
	gorm.Model
	Balance int `gorm:"comment:账户余额"`
}

// 定义transactions表
type transactions struct {
	gorm.Model
	FromAccountID int `gorm:"comment:转出账户ID"`
	ToAccountID   int `gorm:"comment:转入账户ID"`
	Amount        int `gorm:"comment:转账金额"`
}

// 题目3：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 题目4：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     int     `db:"id"`
	Title  string  `gorm:"size:255;comment:图书名称"`
	Author string  `gorm:"size:255;comment:图书作者"`
	Price  float64 `gorm:"comment:图书价格"`
}

// 定义实体几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	gorm.Model
	Name      string `gorm:"size:255;comment:用户姓名"`
	Posts     []Post
	PostCount int `gorm:"default:0;comment:文章数量统计"`
}

type Post struct {
	gorm.Model
	Title       string `gorm:"size:255;comment:文章标题"`
	Content     string `gorm:"size:255;comment:文章内容"`
	UserID      uint
	Comments    []Comment
	CommentStat string `gorm:"size:50;default:'有评论';comment:评论状态"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"size:255;comment:评论内容"`
	PostID  uint
}

// Post 创建前执行钩子：更新 User 的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).
		Error
	return
}

// Comment 删除后执行钩子：检查评论数，更新文章状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	// 查询该文章剩余的评论数
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	// 如果没有评论，更新文章的评论状态
	if count == 0 {
		err = tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_stat", "无评论").
			Error
	}
	return
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gostudy?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 题目1 基本CRUD操作
	// 创建学生表
	//db.AutoMigrate(&student{})
	//// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//result := db.Create(&student{Name: "张三", Age: 20, Grade: "三年级"})
	//if result.Error != nil {
	//	// 插入失败
	//	fmt.Println("插入失败:", result.Error)
	//} else {
	//	fmt.Println("插入成功")
	//}
	//// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//var students []student
	//db.Where("age > ?", 18).Find(&students)
	//fmt.Println("查询 students 表中所有年龄大于 18 岁的学生信息为：", students)
	//// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//result = db.Model(&student{}).Where("name = ?", "张三").Update("grade", "四年级")
	//if result.Error != nil {
	//	// 更新失败
	//	fmt.Println("更新失败:", result.Error)
	//} else {
	//	fmt.Println("更新 students 表中姓名为 '张三' 的学生年级为 '四年级' 成功")
	//}
	//
	//// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。先新增一个李四
	//result = db.Create(&student{Name: "李四", Age: 10, Grade: "三年级"})
	//if result.Error != nil {
	//	// 插入失败
	//	fmt.Println("插入失败:", result.Error)
	//} else {
	//	fmt.Println("插入 students 表中姓名为 '李四' 的学生记录成功")
	//}
	//result = db.Where("age < ?", 15).Delete(&student{})
	//if result.Error != nil {
	//	// 删除失败
	//	fmt.Println("删除失败:", result.Error)
	//} else {
	//	fmt.Println("删除students 表中年龄小于 15 岁的学生记录成功")
	//}

	// 题目2 事务语句
	// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	// 创建account表
	//db.AutoMigrate(&account{})
	//// 创建transactions表
	//db.AutoMigrate(&transactions{})
	//// 新增账户A
	//db.Create(&account{Balance: 1000})
	//// 新增账户B
	//db.Create(&account{Balance: 500})
	//// 转账
	//errRes := db.Transaction(func(tx *gorm.DB) error {
	//	// 扣除账户A余额（带条件，避免并发问题）
	//	res := tx.Model(&account{}).
	//		Where("id = ? AND balance >= ?", 1, 100).
	//		Update("balance", gorm.Expr("balance - ?", 100))
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//	if res.RowsAffected == 0 {
	//		return errors.New("账户A余额不足或不存在")
	//	}
	//
	//	// 给账户B加钱
	//	res = tx.Model(&account{}).
	//		Where("id = ?", 2).
	//		Update("balance", gorm.Expr("balance + ?", 100))
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//	if res.RowsAffected == 0 {
	//		return errors.New("账户B不存在")
	//	}
	//
	//	// 记录转账信息
	//	if err := tx.Create(&transactions{
	//		FromAccountID: 1,
	//		ToAccountID:   2,
	//		Amount:        100,
	//	}).Error; err != nil {
	//		return errors.New("记录转账信息失败")
	//	}
	//
	//	return nil // 返回 nil 表示提交事务
	//})
	//
	//if errRes != nil {
	//	fmt.Println("转账失败:", errRes)
	//} else {
	//	fmt.Println("转账成功")
	//}

	// 题目3：使用SQL扩展库进行查询
	//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	//要求 ：
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

	// 使用sqlx链接数据库
	//dsn := "root:123456@tcp(127.0.0.1:3306)/gostudy?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//// 创建employee表
	//schema := `
	//CREATE TABLE IF NOT EXISTS employees (
	//	id INT AUTO_INCREMENT PRIMARY KEY COMMENT '员工ID',
	//	name VARCHAR(50) NOT NULL COMMENT '员工姓名',
	//	department VARCHAR(50) NOT NULL COMMENT '部门',
	//	salary FLOAT64 NOT NULL COMMENT '工资'
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工表';`
	//_, err = db.Exec(schema)
	//if err != nil {
	//	log.Fatalln("建表失败:", err)
	//}
	//
	//// 2️.插入测试数据
	//db.MustExec("TRUNCATE TABLE employees") // 清空旧数据
	//
	//insertSQL := `
	//INSERT INTO employees (name, department, salary) VALUES
	//('张三', '技术部', 12000.00),
	//('李四', '技术部', 15000.00),
	//('王五', '市场部', 10000.00),
	//('赵六', '人事部', 8000.00),
	//('钱七', '技术部', 20000.00),
	//('孙八', '市场部', 18000.00);`
	//_, err = db.Exec(insertSQL)
	//if err != nil {
	//	log.Fatalln("插入测试数据失败:", err)
	//}
	//
	//// 3.查询技术部员工
	//var techEmployees []Employee
	//err = db.Select(&techEmployees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	//if err != nil {
	//	log.Fatalln("查询技术部员工失败:", err)
	//}
	//fmt.Println("技术部员工:")
	//for _, e := range techEmployees {
	//	fmt.Printf("ID:%d, Name:%s, Dept:%s, Salary:%.2f\n", e.ID, e.Name, e.Department, e.Salary)
	//}
	//
	//// 4.查询工资最高的员工
	//var topEmployee Employee
	//err = db.Get(&topEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	//if err != nil {
	//	log.Fatalln("查询工资最高员工失败:", err)
	//}
	//fmt.Println("工资最高的员工:")
	//fmt.Printf("ID:%d, Name:%s, Dept:%s, Salary:%.2f\n", topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)

	// 题目4：实现类型安全映射
	//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//要求 ：
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	// 创建book表
	//schema := `
	//CREATE TABLE IF NOT EXISTS books(
	//	id INT AUTO_INCREMENT PRIMARY KEY COMMENT '员工ID',
	//	title VARCHAR(50) NOT NULL COMMENT '标题',
	//	author VARCHAR(50) NOT NULL COMMENT '作者',
	//	price DECIMAL(10,2) NOT NULL COMMENT '价格'
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图书表';`
	//_, err = db.Exec(schema)
	//if err != nil {
	//	log.Fatalln("建图书表失败:", err)
	//}
	//// 2️.插入测试数据
	//db.MustExec("TRUNCATE TABLE books") // 清空旧数据
	//
	//insertSQL := `
	//INSERT INTO books (title, author, price) VALUES
	//('Go语言程序设计', '张三', 50.00),
	//('Python从入门到精通', '李四', 60.00),
	//('MySQL数据库原理', '王五', 70.00),
	//('Java核心技术', '赵六', 80.00),
	//('C++程序设计', '钱七', 90.00),
	//('JavaScript高级程序设计', '孙八', 30.00);`
	//_, err = db.Exec(insertSQL)
	//if err != nil {
	//	log.Fatalln("插入测试数据失败:", err)
	//}
	//
	//// 3.查询价格大于50的图书
	//var books []Book
	//err = db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	//if err != nil {
	//	log.Fatalln("查询价格大于50的图书失败:", err)
	//}
	//fmt.Println("价格大于50的图书:")
	//for _, b := range books {
	//	fmt.Printf("ID:%d, Title:%s, Author:%s, Price:%.2f\n", b.ID, b.Title, b.Author, b.Price)
	//}

	// gorm进阶
	// 题目1：模型定义
	//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	//要求 ：
	//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//题目2：关联查询
	//基于上述博客系统的模型定义。
	//要求 ：
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	// 插入测试数据
	// 插入测试数据
	//db.Create(&User{Name: "张三"})
	//db.Create(&User{Name: "李四"})
	//db.Create(&Post{Title: "文章1", UserID: 1})
	//db.Create(&Post{Title: "文章2", UserID: 1})
	//db.Create(&Post{Title: "文章3", UserID: 2})
	//db.Create(&Comment{Content: "评论1", PostID: 1})
	//db.Create(&Comment{Content: "评论2", PostID: 1})
	//db.Create(&Comment{Content: "评论3", PostID: 2})
	//db.Create(&Comment{Content: "评论4", PostID: 2})
	//db.Create(&Comment{Content: "评论5", PostID: 3})
	//// 1. 查询某个用户发布的所有文章及评论
	//var posts []Post
	//db.Preload("Comments").Joins("User").Where("users.name = ?", "张三").Find(&posts)
	//for _, p := range posts {
	//	fmt.Printf("文章标题:%s\n", p.Title)
	//	for _, c := range p.Comments {
	//		fmt.Printf("  评论内容:%s\n", c.Content)
	//	}
	//}
	//
	//// 2. 查询评论数量最多的文章
	//var post Post
	//db.Preload("Comments").
	//	Joins("left join comments on comments.post_id = posts.id").
	//	Group("posts.id").
	//	Order("count(comments.id) desc").
	//	Limit(1).
	//	Find(&post)
	//
	//fmt.Printf("评论数量最多的文章标题:%s, 评论数量:%d\n", post.Title, len(post.Comments))

	// 题目3：钩子函数
	//继续使用博客系统的模型。
	//要求 ：
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// 插入数据测试
	u := User{Name: "田七"}
	db.Create(&u)

	p := Post{Title: "文章7", UserID: u.ID}
	db.Create(&p) // 会触发 AfterCreate 钩子，更新 User.PostCount

	c1 := Comment{Content: "评论7", PostID: p.ID}
	db.Create(&c1)

	c2 := Comment{Content: "评论8", PostID: p.ID}
	db.Create(&c2)

	db.Delete(&c1)
	db.Delete(&c2) // 会触发 AfterDelete 钩子，Post.CommentStat 更新为 "无评论"
}
