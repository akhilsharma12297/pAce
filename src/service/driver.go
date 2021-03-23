package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var restPort = "8080"

func DriverMethod() {

	fmt.Println("Server Started!")
	fmt.Println("Listening to port", restPort)

	router := gin.Default()

	v1 := router.Group("/transactionservice")
	{
		v1.PUT("/transaction/:transaction_id", PutTransaction)
		v1.GET("/transaction/:transaction_id", GetTransaction)
		v1.GET("/types/:types_value", GetTransactionType)
		v1.GET("/sum/:transaction_id", GetTransactionSum)

	}

	log.Fatal(router.Run(":" + restPort))

}
