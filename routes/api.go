package routes

import (
	"bribrain-initial/controller"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Router() {

	defer func() {
		if err := recover(); err != nil {
			logrus.Warning("Panic Detected : " + fmt.Sprint(err))
		}
	}()
	rkaController := new(controller.RkaController)
	routing := echo.New()
	routing.POST("/api/adds/users", rkaController.Adduser)
	routing.POST("/api/adds/rka", rkaController.AddRKA)
	routing.POST("/api/adds/users/rka", rkaController.AddUserRKA)
	routing.GET("/api/accumulation/rka", rkaController.AccumulationRKA)

	routing.Logger.Fatal(routing.Start(":" + os.Getenv("APP_PORT")))
}
