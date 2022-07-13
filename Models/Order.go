package Models

import (
	"Day4_5/Config"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// create product
func CreateOrder(order *Order) (err error) {
	var product Product
	if err = Config.DB.Where("id = ?", order.ProductID).First(&product).Error; err != nil {
		return err
	}
	if product.Quantity < order.Quantity {
		return errors.New("Product quantiy is not sufficient to execute order")
	} else {
		product.Quantity -= order.Quantity
	}
	if err = Config.DB.Save(&product).Error; err != nil {
		return err
	}
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// get product by Id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

// get all orders
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}
