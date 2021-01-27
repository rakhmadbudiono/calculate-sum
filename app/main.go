package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	
	_calculationHttpDelivery "github.com/rakhmadbudiono/calculate-sum/calculation/delivery/http"
	_calculationHttpDeliveryMiddleware "github.com/rakhmadbudiono/calculate-sum/calculation/delivery/http/middleware"
	_calculationUcase "github.com/rakhmadbudiono/calculate-sum/calculation/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	middL := _calculationHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	c := _calculationUcase.NewCalculationUsecase(timeoutContext)
	_calculationHttpDelivery.NewCalculationHandler(e, c)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
