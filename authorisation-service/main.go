package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prithuadhikary/let-us-go/authorisation-service/controller"
	"github.com/prithuadhikary/let-us-go/authorisation-service/service"
	"github.com/prithuadhikary/let-us-go/common/auth"
	"github.com/prithuadhikary/let-us-go/common/client"
	"log"
)

func main() {
	signingKey := []byte("eternalSecret")
	userService := client.NewUserServiceClient()
	tokenService := auth.NewTokenService(signingKey)
	loginService := service.NewLoginService(userService, tokenService)

	engine := gin.Default()

	controller.NewLoginController(engine, loginService)

	log.Fatal(engine.Run(":8081"))
}
