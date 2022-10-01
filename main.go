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

	userDal := dal.NewUserDal(db)

	// changed this line from origin
	userService := service.NewUserService(
		*userDal,
	)

	log.Printf("userService: %v", userService)
}
