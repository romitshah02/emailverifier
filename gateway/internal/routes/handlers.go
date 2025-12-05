package routes

import (
	"gateway/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/routes", listRoutes)
		admin.POST("/routes", createRoute)
		admin.PUT("/routes/:id", updateRoute)
		admin.GET("/routes/:id", getRoute)
		admin.DELETE("/routes/:id", deleteRoute)
	}
}

func listRoutes(c *gin.Context) {
	routes, err := ListRoutes()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"routes": routes})
}

func createRoute(c *gin.Context) {
	var route models.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := CreateRoute(&route); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Route Created", "route": route})
}

func deleteRoute(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := DeleteRoute(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Route Deleted"})

}

func getRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	route, err := GetRouteID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"route": route})
}

func updateRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var route models.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	route.ID = uint(id)

	if err := UpdateRoute(&route); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Route updated", "route": route})
}
