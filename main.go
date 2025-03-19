package main

import (
	"century/controllers"
	"century/routes"
	"century/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()
	service := services.NewBankService()
	controller := controllers.NewBankController(service)
	routes.AddMainAPIs(engine, controller)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    1200 * time.Second,
		WriteTimeout:   1200 * time.Second,
		MaxHeaderBytes: 5 * http.DefaultMaxHeaderBytes,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("Server Stopped:", err)
	}
	fmt.Println("Server started on :8080")

}
