package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	// inicializar sqlite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqli > %v", err)
	}

	//nil es null em go
	//testando o mensagem de error
	// return errors.New("fake fake error")
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// inicia logger
	logger = NewLogger(p)
	return logger
}
