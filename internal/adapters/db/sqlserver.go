package db

import (
	_ "MQTTHub/internal/core/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// const (
// 	server   = "localhost"
// 	port     = 1433
// 	user     = "SA"
// 	password = "Test1234"
// 	database = "TitDB"
// )

// // github.com/denisenkom/go-mssqldb
// var dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
// 	server, user, password, port, database)

type Adapter struct {
	db *gorm.DB
}

// NewAdapter creates a new Adapter
func NewAdapter(connectionString string) (*Adapter, error) {
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

// Create takes a pointer of a model struct and inserts it to db
func (da Adapter) Create(modelP interface{}) error {
	err := da.db.Create(modelP).Error
	if err != nil {
		return err
	}

	return nil
}

// Find performs a query by sending a property, a value to match and a pointer to the model of the table
func (da Adapter) Find(parameter string, value string, modelP interface{}) error {

	err := da.db.Where(fmt.Sprintf("%s = ?", parameter), value).First(modelP).Error
	if err != nil {
		return err
	}

	return nil
}
