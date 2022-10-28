package gorm

import (
	model "base-frame/services/user/model"
	query "base-frame/services/user/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	ctx:=context.Background()
	u := query.Use(db).User
	user:=model.User{
		Name:"test",
		CreateTime: time.Now(),
	}
	u.WithContext(ctx).Select(u.Name, u.CreateTime).Create(&user)
	// INSERT INTO `users` (`name`,`create_time`) VALUES ("modi", 18)

	allUsers, _ := u.WithContext(ctx).Find()
	// SELECT * FROM users;

	singleUser, _ := u.WithContext(ctx).Where(u.Name.Eq("modi")).First()
	// SELECT * FROM users WHERE name = 'modi' ORDER BY id LIMIT 1;

	someUsers, _ := u.WithContext(ctx).Where(u.Name.Eq("modi"), u.ID.Gte(17)).Find()
	// SELECT * FROM users WHERE name = 'modi' AND id >= 17;
	fmt.Println(allUsers, singleUser, someUsers)
}