package controller

import (
	"bribrain-initial/model"
	"bribrain-initial/repository"
	"bribrain-initial/service"
	"strconv"

	"bribrain-initial/rpc/rka"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type RkaController struct {
	repository repository.Services
}

func NewControllerRka() *RkaController {
	repo := service.NewServiceRka()
	return &RkaController{
		repository: repo,
	}
}

func (*RkaController) Adduser(c echo.Context) (err error) {
	request := new(rka.User)

	if err := c.Bind(request); err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}

	services := service.NewServiceRka()
	response, err := services.Adduser(c.Request().Context(), request)
	if err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}

	if response.Code != 200 {
		logrus.Error("call rpc func failed : " + response.Message.(string))
		return model.ResponseContext(int(response.Code), response.Message, nil, c)
	}

	return model.ResponseContext(200, response.Message, response.Data, c)
}

func (*RkaController) AddRKA(c echo.Context) (err error) {
	request := new(rka.RKA)

	if err := c.Bind(request); err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}
	services := service.NewServiceRka()
	response, err := services.AddRKA(c.Request().Context(), request)
	if err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}

	if response.Code != 200 {
		logrus.Error("call rpc func failed : " + response.Message.(string))
		return model.ResponseContext(int(response.Code), response.Message, nil, c)
	}

	return model.ResponseContext(200, response.Message, response.Data, c)
}

func (*RkaController) AddUserRKA(c echo.Context) (err error) {
	request := new(rka.RequestUserRKA)

	if err := c.Bind(request); err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}
	services := service.NewServiceRka()
	response, err := services.AddUserRKA(c.Request().Context(), request)
	if err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}

	if response.Code != 200 {
		logrus.Error("call rpc func failed : " + response.Message.(string))
		return model.ResponseContext(int(response.Code), response.Message, nil, c)
	}

	return model.ResponseContext(200, response.Message, response.Data, c)
}

func (*RkaController) AccumulationRKA(c echo.Context) (err error) {
	userId := c.QueryParam("user_id")
	year := c.QueryParam("year")
	month := c.QueryParam("month")
	userIDInt, _ := strconv.Atoi(userId)
	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)

	services := service.NewServiceRka()
	response, err := services.AccumulationRKA(c.Request().Context(), &rka.RequestAccumulation{
		UserId: int64(userIDInt),
		Yaer:   int32(yearInt),
		Month:  int32(monthInt),
	})
	if err != nil {
		return model.ResponseContext(400, err.Error(), nil, c)
	}

	if response.Code != 200 {
		logrus.Error("call rpc func failed : " + response.Message.(string))
		return model.ResponseContext(int(response.Code), response.Message, nil, c)
	}

	return model.ResponseContext(200, response.Message, response.Data, c)
}
