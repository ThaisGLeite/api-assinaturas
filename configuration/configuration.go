package configuration

import (
	"log"
	"time"
)

type Logfile struct {
	ErrorLogger log.Logger
	InfoLogger  log.Logger
}

func Check(erro error, loggin Logfile) {
	if erro != nil {
		loggin.ErrorLogger.Fatal(erro)
	}
}

func CheckValidadeAssinante(data string, loggin Logfile) string {
	status := "valido"

	date, err := time.Parse("2006-01-02", data)
	Check(err, loggin)

	agora := time.Now().UTC()

	if date.Before(agora) {
		status = "invalido"
	}

	return status
}
