package controller

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"order-management/config"
	"order-management/shared/discovery"
)

type GatewayController struct {
	productProxy *httputil.ReverseProxy
	orderProxy   *httputil.ReverseProxy
	authProxy    *httputil.ReverseProxy
}

func NewGatewayController() (*GatewayController, error) {
	consulClient, err := discovery.NewConsulClient()
	if err != nil {
		log.Printf("Consul client unavailable; Gateway will use configured fallback URLs: %v", err)
		consulClient = nil
	}
	return &GatewayController{
		productProxy: newLoadBalancedProxy(consulClient, config.Get("PRODUCT_SERVICE_NAME", "product-service"), config.Get("PRODUCT_SERVICE_URL", "http://localhost:8081")),
		orderProxy:   newLoadBalancedProxy(consulClient, config.Get("ORDER_SERVICE_NAME", "order-service"), config.Get("ORDER_SERVICE_URL", "http://localhost:8082")),
		authProxy:    newLoadBalancedProxy(consulClient, config.Get("AUTH_SERVICE_NAME", "auth-service"), config.Get("AUTH_SERVICE_URL", "http://localhost:8084")),
	}, nil
}

func newLoadBalancedProxy(consulClient *consulapi.Client, serviceName, fallback string) *httputil.ReverseProxy {
	balancer := discovery.NewRoundRobin(consulClient, serviceName, fallback)
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			target, err := balancer.Next()
			if err != nil {
				target, _ = url.Parse(fallback)
			}
			if target == nil {
				target, _ = url.Parse(fallback)
			}
			targetQuery := target.RawQuery
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
			if targetQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = targetQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
			}
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, "upstream service unavailable", http.StatusBadGateway)
		},
	}
}

func (c *GatewayController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "api-gateway"})
}
func (c *GatewayController) Product(ctx *gin.Context) {
	c.productProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
func (c *GatewayController) Order(ctx *gin.Context) { c.orderProxy.ServeHTTP(ctx.Writer, ctx.Request) }
func (c *GatewayController) Auth(ctx *gin.Context)  { c.authProxy.ServeHTTP(ctx.Writer, ctx.Request) }
