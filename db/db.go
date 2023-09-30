package db

import (
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Adapter *gormadapter.Adapter
var Enforcer *casbin.Enforcer

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		panic("Failed to connect to DB")
	}
}

func CasbinAdapter() {
	var err error
	Adapter, err = gormadapter.NewAdapterByDB(DB)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}
}

func CasbinEnforcer() {
	var err error
	Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", Adapter)
	if err != nil {
		panic(fmt.Errorf("failed to create casbin enforcer %w", err))
	}
}
