package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authHandler(c *gin.Context) {
	proxyRequest(c, "http://auth-service:8080")
}

func orderHandler(c *gin.Context) {
	proxyRequest(c, "http://order-service:8081")
}

func proxyRequest(c *gin.Context, serviceURL string) {
	req, err := http.NewRequest(c.Request.Method, serviceURL+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

func main() {
	r := gin.Default()

	r.Any("/auth/*any", authHandler)
	r.Any("/orders/*any", orderHandler)

	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Failed to run API Gateway: %v", err)
	}
}
