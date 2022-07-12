package Controllers

import (
	"Day4_5/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create Product
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		order.Status = "order placed"
		c.JSON(http.StatusOK, order)
	}
}

// getproduct by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
