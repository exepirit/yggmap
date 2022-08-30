package main

import (
	"net/http"

	v1 "github.com/exepirit/yggmap/internal/api/v1"
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	Route   string
	Method  string
	Handler gin.HandlerFunc
}

var Endpoints = []Endpoint{
	{
		Route:   "/api/v1/network",
		Method:  http.MethodGet,
		Handler: GetStaticNetworkMap,
	},
}

// GetStaticNetwork function purpose is provide static network map for testing.
func GetStaticNetworkMap(ctx *gin.Context) {
	network := v1.NetworkDto{
		Nodes: []v1.NodeDto{
			{
				PublicKey: "baa97089f3973fe68c5b7249f45ebde91aa266f4f22682b6b2840677ae654938",
				AdditionalInfo: map[string]interface{}{
					"os": "linux",
				},
			},
			{
				PublicKey: "ee88dd4a5c1a7261ba4492690b701ccbb6a86452d40b19352d250082bcf58f86",
				AdditionalInfo: map[string]interface{}{
					"os": "windows",
				},
			},
			{
				PublicKey: "c9d88e46341c357111a0ef8b216c1d3bf8d363a63b20e7f7f851ecc960c7408b",
				AdditionalInfo: map[string]interface{}{
					"os": "android",
				},
			},
		},
		Edges: []v1.EdgeDto{
			{
				From: "baa97089f3973fe68c5b7249f45ebde91aa266f4f22682b6b2840677ae654938",
				To:   "ee88dd4a5c1a7261ba4492690b701ccbb6a86452d40b19352d250082bcf58f86",
			},
			{
				From: "ee88dd4a5c1a7261ba4492690b701ccbb6a86452d40b19352d250082bcf58f86",
				To:   "c9d88e46341c357111a0ef8b216c1d3bf8d363a63b20e7f7f851ecc960c7408b",
			},
			{
				From: "c9d88e46341c357111a0ef8b216c1d3bf8d363a63b20e7f7f851ecc960c7408b",
				To:   "baa97089f3973fe68c5b7249f45ebde91aa266f4f22682b6b2840677ae654938",
			},
		},
	}

	ctx.JSON(http.StatusOK, network)
}
