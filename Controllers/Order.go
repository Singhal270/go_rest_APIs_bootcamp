package Controllers

import (
	"Day4_5/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllorder
func GetAllOrders(c *gin.Context) {
	var orders []Models.Order
	err := Models.GetAllOrders(&orders)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

//Create order
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	order.Status = "order placed"
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// getorder by id
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
