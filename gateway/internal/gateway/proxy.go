package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(c *gin.Context) {
	method := c.Request.Method
	path := c.Request.URL.Path

	if after, ok := strings.CutPrefix(path, "/proxy"); ok {
		path = after
	}

	route := GetRoute(method, path)
	if route == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No route found"})
		return
	}

	log.Printf("✅ Proxying %s %s -> %s", method, path, route.UpstreamURL)

	targetURL, err := url.Parse(route.UpstreamURL)
	if err != nil {
		log.Printf("❌ Invalid URL: %v", err)
		c.JSON(500, gin.H{"error": "Invalid upstream URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// CRITICAL: Modify request for proxy
	c.Request.URL.Host = targetURL.Host
	c.Request.URL.Scheme = targetURL.Scheme
	c.Request.Host = targetURL.Host

	proxy.ServeHTTP(c.Writer, c.Request)
	log.Printf("✅ Proxy completed")
}
