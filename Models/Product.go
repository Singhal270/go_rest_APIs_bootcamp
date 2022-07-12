package Models

import (
	"Day4_5/Config"

	_ "github.com/go-sql-driver/mysql"
)

// get all products
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

// create product
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// get product by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//Update product
func UpdateProduct(product *Product, id string) (err error) {
	if err = Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(product *Product, id string) (err error) {

	if err = Config.DB.Where("id = ?", id).Delete(product).Error; err != nil {
		return err
	}
	return nil
}
