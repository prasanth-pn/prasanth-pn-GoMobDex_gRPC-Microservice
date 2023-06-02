package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/prasanth-pn/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
