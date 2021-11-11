package  github.com/dm0610/go-lvl2-hw3-01

import (
	"log"
)

func Config() {
	appHost := "localhost"
	appPort := "8080"
	log.Printf("App start on %s:%s address!\n", appHost, appPort)
}
