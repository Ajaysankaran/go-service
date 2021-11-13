package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/theAimOne/service-gateway/config"
	"github.com/theAimOne/service-gateway/service"
	"github.com/theAimOne/service-gateway/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	hostname     = "localhost"
	host_port    = 5432
	username     = "postgres"
	password     = "postgres"
	databasename = "postgres"
	schemaname   = "servpoint"
)

var err error

func SetupDatabase() {
	pg_con_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostname, username, password, databasename, host_port)
	config.DB, err = gorm.Open(postgres.Open(pg_con_string), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := config.DB.DB()
	sqlDB.Exec(fmt.Sprintf("set search_path=%s", schemaname))

	config.DB.AutoMigrate(&user.User{})
	config.DB.AutoMigrate(&service.Service{})
}
