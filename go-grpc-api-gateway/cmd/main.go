package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/go-grpc-api-gateway/pkg/auth"
	"github.com/prasanth-pn/go-grpc-api-gateway/pkg/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println(authSvc)

	r.Run(c.Port)
}
