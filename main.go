package main

import (
	"bribrain-initial/routes"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
	}

	routes.Router() //Initialization routes http
}
