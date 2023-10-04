package main

import (
	"github.com/halrrik/gopportunities/config"
	"github.com/halrrik/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		// panic(err)
		// fmt.Println(err)
		logger.Errorf("config initializa error: %v", err)
		return
	}

	//initialize router
	//SEMPRE QUE TEM UMA VARIAVEL QUE INICIA CON
	//MAYS ES UNA VARIAVEL DE UN PACOTE QUE ESTA SIENDO EXPORTADA
	router.Initialize()
}
