package Routes

import (
	"Day4_5/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("")
	{
		grp1.GET("products", Controllers.GetAllProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)

		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.GetOrderByID)

	}
	return r
}
