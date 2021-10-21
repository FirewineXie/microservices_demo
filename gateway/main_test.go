package main

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/gateway/api/v1"
	"net/http"
	"testing"
)

/*
@Time : 2021/10/21 22:59
@Author : Firewine
@File : main_test.go
@Software: GoLand
@Description:
*/

func Test_main(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		product, err := getProduct(c.Request.Context(), "OLJCESPC7Z")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"messasge": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
			"data":product,
		})
	})
	r.Run()
}
func getProduct(ctx context.Context, id string) (*v1.Product, error) {
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9007", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewProductCatalogServiceClient(conn).GetProduct(context.Background(), &v1.GetProductRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil

}
