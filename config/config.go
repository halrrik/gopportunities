package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	//nil es null em go
	//testando o mensagem de error
	// return errors.New("fake fake error")
	return nil
}

func GetLogger(p string) *Logger {
	// inicia logger
	logger = NewLogger(p)
	return logger
}
