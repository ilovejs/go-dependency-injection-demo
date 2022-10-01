package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"DIdemo/dal"
	"DIdemo/service"
)

func main() {
	dsn := "host=localhost user=di_demo password=di_111 dbname=di_user port=5432 sslmode=disable TimeZone=Australia/Sydney"

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	// 通过将 db 实例对象注入到 dal 中，再将 dal 实例对象注入到 service 中，
	// 实现了层级间的依赖注入。解耦了部分依赖关系
	userDal := dal.NewUserDal(db)

	// changed this line from origin
	userService := service.NewUserService(
		*userDal,
	)

	log.Printf("userService: %v", userService)
}
