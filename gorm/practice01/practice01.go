package practice01

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Author struct {
	Name  string
	Email string
}

type Blog1 struct {
	Author
	ID      int
	Upvotes int32
}

// equals
type Blog2 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Run(db *gorm.DB) {
	// 自动创建表
	// db.AutoMigrate(&User{})

	// 写入数据
	db.Create(&User{Name: "张三", Age: 18})
	// 读取数据
	var user User
	db.First(&user, "name = ?", "张三")
	fmt.Println("用户id：", user.ID)
	fmt.Println("用户姓名：", user.Name)
	fmt.Println("用户年龄：", user.Age)

	// 更新数据
	db.Model(&user).Update("Name", "李四")
	// 读取更新后的数据
	db.First(&user, "name = ?", "李四")
	fmt.Println("用户id：", user.ID)
	fmt.Println("用户姓名：", user.Name)
	fmt.Println("用户年龄：", user.Age)

	// 删除数据
	db.Delete(&user, "name = ?", "李四")
	var deletedUser User
	err := db.First(&deletedUser, "name = ?", "李四").Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("用户已被删除")
	} else {
		fmt.Println("用户未被删除，用户名：", deletedUser.Name)
	}
}
